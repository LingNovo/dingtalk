package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	pb "github.com/LingNovo/dingtalk/protos/oapi"
)

type Unmarshallable interface {
	CheckError() error
}

func (dtc *OapiSrv) httpRPC(path string, params url.Values, requestData interface{}, responseData Unmarshallable, isvGetCInfo ...interface{}) error {
	if dtc.DevType == "company" {
		if dtc.AccessToken != "" {
			if params == nil {
				params = url.Values{}
			}
			if params.Get("access_token") == "" {
				params.Set("access_token", dtc.AccessToken)
			}
		}
	}
	if dtc.DevType == "isv" {
		cur := isvGetCInfo[0]
		switch v := cur.(type) {
		case *pb.DTIsvGetCompanyInfo:
			switch path {
			case "service/get_permanent_code", "service/activate_suite", "service/get_corp_token", "service/get_auth_info":
				if dtc.SuiteAccessToken != "" {
					if params == nil {
						params = url.Values{}
					}
					if params.Get("suite_access_token") == "" {
						params.Set("suite_access_token", dtc.SuiteAccessToken)
					}
				}
			default:
				if v.AuthAccessToken != "" {
					if params == nil {
						params = url.Values{}
					}
					if params.Get("access_token") == "" {
						params.Set("access_token", v.AuthAccessToken)
					}
				}
			}
		default:
			panic(errors.New("ERROR: *DTIsvGetCompanyInfo Error"))
		}
	}
	if dtc.DevType == "personalMini" {
		if dtc.SNSAccessToken != "" && path != "sns/getuserinfo" {
			if params == nil {
				params = url.Values{}
			}
			if params.Get("access_token") == "" {
				params.Set("access_token", dtc.SNSAccessToken)
			}
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

func (dtc *OapiSrv) httpSNS(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	if dtc.SNSAccessToken != "" && path != "sns/getuserinfo" {
		if params == nil {
			params = url.Values{}
		}
		if params.Get("access_token") == "" {
			params.Set("access_token", dtc.SNSAccessToken)
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

func (dtc *OapiSrv) httpSSO(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	if dtc.SSOAccessToken != "" {
		if params == nil {
			params = url.Values{}
		}
		if params.Get("access_token") == "" {
			params.Set("access_token", dtc.SSOAccessToken)
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

type TopMapRequest map[string]interface{}

func (t TopMapRequest) keys() []string {
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

func (dtc *OapiSrv) httpTOP(requestData interface{}, responseData interface{}) error {
	var params []string
	var paramsJoin string
	var cipher []byte
	var cipherString string
	if body, ok := requestData.(TopMapRequest); ok {
		body["sign_method"] = dtc.TopConfig.TopSignMethod
		if dtc.DevType == "company" {
			body["session"] = dtc.AccessToken
		}
		body["format"] = dtc.TopConfig.TopFormat
		body["v"] = dtc.TopConfig.TopVersion
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		body["timestamp"] = tm.Format("2006-01-02 03:04:05 PM")
		if dtc.TopConfig.TopFormat == "json" {
			body["simplify"] = dtc.TopConfig.TopSimplify
		}
		params = sortParamsKey(body)
		paramsJoin = strings.Join(params, "")
		if dtc.TopConfig.TopSignMethod == signMD5 {
			paramsJoin = dtc.TopConfig.TopSecret + paramsJoin + dtc.TopConfig.TopSecret
			cipher = encryptMD5(paramsJoin)
		}
		if dtc.TopConfig.TopSignMethod == signHMAC {
			cipher = encryptHMAC(paramsJoin, dtc.TopConfig.TopSecret)
		}
		cipherString = byteToHex(cipher)
		body["sign"] = cipherString
		fmt.Printf("Top Params=%v\n", body)
		return dtc.httpRequest("tapi", nil, addPostBody(body), nil, responseData)
	}
	return errors.New("requestData Not TopMapRequest Type")
}

func addPostBody(topMap TopMapRequest) url.Values {
	body := url.Values{}
	for k, v := range topMap {
		switch v.(type) {
		case []string:
			for _, h := range v.([]string) {
				body.Add(k, h)
			}
		case []int:
			for _, h := range v.([]int) {
				body.Add(k, string(h))
			}
		default:
			body.Add(k, fmt.Sprintf("%v", v))
		}
	}
	return body
}

func encryptMD5(paramsJoin string) []byte {
	hMd5 := md5.New()
	hMd5.Write([]byte(paramsJoin))
	return hMd5.Sum(nil)
}

func encryptHMAC(paramsJoin string, secret string) []byte {
	hHmac := hmac.New(md5.New, []byte(secret))
	hHmac.Write([]byte(paramsJoin))
	return hHmac.Sum([]byte(""))
}

func byteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(strings.ToUpper(s))
	}
	return buffer.String()
}

func sortParamsKey(topParams TopMapRequest) []string {
	var t []string
	keys := topParams.keys()
	sort.Strings(keys)
	for _, k := range keys {
		t = append(t, k+fmt.Sprintf("%s", topParams[k]))
	}
	return t
}

func (dtc *OapiSrv) httpRequest(tagType string, path interface{}, params url.Values, requestData interface{}, responseData interface{}) error {
	var request *http.Request
	var requestUrl string
	client := dtc.HTTPClient

	if tagType == "oapi" {
		requestUrl = OAPIURL + path.(string) + "?" + params.Encode()
		fmt.Printf("requestUrl=%s\n", requestUrl)
		if requestData != nil {
			d, _ := json.Marshal(requestData)
			request, _ = http.NewRequest("POST", requestUrl, bytes.NewReader(d))
			request.Header.Set("Content-Type", typeJSON+"; charset=UTF-8")
		} else {
			request, _ = http.NewRequest("GET", requestUrl, nil)
		}
	}
	if tagType == "tapi" {
		requestUrl = TOPAPIURL
		request, _ = http.NewRequest("POST", requestUrl, strings.NewReader(params.Encode()))
		request.Header.Set("Content-Type", typeForm+"; charset=UTF-8")
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Server Error: " + resp.Status)
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	pos := len(typeJSON)

	if tagType == "oapi" {
		if len(contentType) >= pos && contentType[0:pos] == typeJSON {
			if content, err := ioutil.ReadAll(resp.Body); err == nil {
				fmt.Println(string(content))
				if err := json.Unmarshal(content, responseData); err != nil {
					return err
				}
				switch responseData.(type) {
				case Unmarshallable:
					resData := responseData.(Unmarshallable)
					return resData.CheckError()
				default:
					v := reflect.ValueOf(responseData)
					v = v.Elem()
					v.SetBytes(content)
				}
			}
		}
	}
	if tagType == "tapi" {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		if dtc.TopConfig.TopFormat == "json" {
			if err := json.Unmarshal(content, responseData); err != nil {
				return err
			}
		} else {
			if err := xml.Unmarshal(content, responseData); err != nil {
				return err
			}
		}
		switch responseData.(type) {
		case Unmarshallable:
			resData := responseData.(Unmarshallable)
			return resData.CheckError()
		default:
			v := reflect.ValueOf(responseData)
			v = v.Elem()
			v.SetBytes(content)
		}

	}
	return nil
}
