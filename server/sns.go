package server

import (
	"fmt"
	"net/url"

	pb "github.com/LingNovo/dingtalk/protos/oapi"
	context "golang.org/x/net/context"
)

// 获取用户授权的持久授权码
func (dtc *OapiSrv) SnsGetPersistentCode(ctx context.Context, req *pb.SnsGetPersistentCodeRequest) (*pb.SnsGetPersistentCodeResponse, error) {
	var data pb.SnsGetPersistentCodeResponse
	requestData := map[string]string{
		"tmp_auth_code": req.Code, //用户授权给钉钉开放应用的临时授权码
	}
	if err := dtc.httpSNS("sns/get_persistent_code", nil, requestData, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取用户授权的SNS_TOKEN
func (dtc *OapiSrv) SnsGetSnsToken(ctx context.Context, req *pb.SnsGetSnsTokenRequest) (*pb.SnsGetSnsTokenResponse, error) {
	var data pb.SnsGetSnsTokenResponse
	requestData := map[string]string{
		"openid":          req.Openid,         //用户的openid
		"persistent_code": req.PersistentCode, //用户授权给钉钉开放应用的持久授权码
	}
	if err := dtc.httpSNS("sns/get_sns_token", nil, requestData, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取用户授权的个人信息
func (dtc *OapiSrv) SnsGetUserInfo(ctx context.Context, req *pb.SnsGetUserInfoRequest) (*pb.SnsGetUserInfoResponse, error) {
	var data pb.SnsGetUserInfoResponse
	params := url.Values{}
	params.Add("sns_token", req.SnsToken) //用户授权给开放应用的token
	if err := dtc.httpSNS("sns/getuserinfo", params, nil, &data); err != nil {
		return nil, err

	}
	return &data, nil
}

// 第三方网站使用钉钉账号登陆
func (dtc *OapiSrv) SnsAuthorize(ctx context.Context, req *pb.SnsAuthorizeRequest) (*pb.SnsAuthorizeResponse, error) {
	resp := &pb.SnsAuthorizeResponse{}
	resp.TargetUrl = fmt.Sprintf("https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=%s&response_type=code&scope=%s&state=%s&redirect_uri=%s",
		dtc.DTConfig.SNSAppId, req.Scope, req.State, req.RedirectUri)
	fmt.Println(resp.TargetUrl)
	return resp, nil
}

// 第三方网站使用钉钉扫码登陆
func (dtc *OapiSrv) Qrconnect(ctx context.Context, req *pb.QrconnectRequest) (*pb.QrconnectResponse, error) {
	resp := &pb.QrconnectResponse{}
	resp.TargetUrl = fmt.Sprintf("https://oapi.dingtalk.com/connect/qrconnect?appid=%s&response_type=code&scope=snsapi_login&state=%s&redirect_uri=%s",
		dtc.DTConfig.SNSAppId, req.State, req.RedirectUri)
	fmt.Println(resp.TargetUrl)
	return resp, nil
}
