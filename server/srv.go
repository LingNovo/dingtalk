package server

import (
	"net/http"
	"sync"
	"time"
)

//oapi rpc server imp
type OapiSrv struct {
	DTConfig              *DTConfig
	TopConfig             *TopConfig
	HTTPClient            *http.Client
	AccessToken           string
	SSOAccessToken        string
	SNSAccessToken        string
	SuiteAccessToken      string
	AccessTokenCache      Cache
	TicketCache           Cache
	SSOAccessTokenCache   Cache
	SNSAccessTokenCache   Cache
	SuiteAccessTokenCache Cache
	DevType               string
	Locker                *sync.Mutex
}

type TopConfig struct {
	TopFormat     string //响应格式。默认为xml格式，可选值：xml，json。
	TopVersion    string
	TopSimplify   bool   //是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false。
	TopSecret     string //签名密钥
	TopSignMethod string //签名方法，signMD5，signHMAC
}

type DTConfig struct {
	TopConfig
	CorpId        string //CorpID是企业在钉钉中的标识；每个企业拥有一个唯一的CorpID
	CorpSecret    string //CorpSecret是企业每个微应用的凭证密钥。
	AgentId       string //代表应用和企业映射关系的ID（appId的实例化ID），同一个ISV应用在不同企业的agentId不一致(编辑微应用时可见)
	SuiteKey      string //套件key，开发者后台创建套件后获取
	SuiteSecret   string //套件secret，开发者后台创建套件后获取
	SuiteTicket   string //钉钉推送的ticket
	ChannelSecret string //服务窗密钥
	SSOSecret     string //微应用管理密钥
	SNSAppId      string //应用ID，同一个ISV应用在不同企业的appId一致,由钉钉开放平台提供给开放应用的唯一标识
	SNSSecret     string //由钉钉开放平台提供的密钥
}

func NewOapiSrv(devType string, config *DTConfig) *OapiSrv {
	c := &OapiSrv{
		DTConfig: &DTConfig{},
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		TopConfig: &TopConfig{
			TopFormat:     topFormat,
			TopSecret:     topSecret,
			TopSignMethod: topSignMethod,
			TopSimplify:   topSimplify,
			TopVersion:    topVersion,
		},
		AccessTokenCache:      NewFileCache("." + devType + "_access_token_file"),
		TicketCache:           NewFileCache("." + devType + "_ticket_file"),
		SSOAccessTokenCache:   NewFileCache("." + devType + "_sso_acess_token_file"),
		SNSAccessTokenCache:   NewFileCache("." + devType + "_sns_access_token_file"),
		SuiteAccessTokenCache: NewFileCache("." + devType + "_suite_access_token_file"),
		Locker:                new(sync.Mutex),
		DevType:               devType,
	}
	if config != nil {
		if config.TopFormat != "" {
			c.TopConfig.TopFormat = config.TopFormat
		}
		if config.TopVersion != "" {
			c.TopConfig.TopVersion = config.TopVersion
		}
		if config.TopSecret != "" {
			c.TopConfig.TopSecret = config.TopSecret
		}
		if config.TopSignMethod != "" {
			c.TopConfig.TopSignMethod = config.TopSignMethod
		}
		if config.TopSimplify {
			c.TopConfig.TopSimplify = config.TopSimplify
		}
		c.DTConfig.CorpId = config.CorpId
		c.DTConfig.AgentId = config.AgentId
		c.DTConfig.CorpSecret = config.CorpSecret
		c.DTConfig.SSOSecret = config.SSOSecret
		c.DTConfig.ChannelSecret = config.ChannelSecret
		c.DTConfig.SNSAppId = config.SNSAppId
		c.DTConfig.SNSSecret = config.SNSSecret
		c.DTConfig.SuiteKey = config.SuiteKey
		c.DTConfig.SuiteSecret = config.SuiteSecret
		c.DTConfig.SuiteTicket = config.SuiteTicket
	}
	return c
}

// 应用服务商
func NewISVOapiSrv(config *DTConfig) *OapiSrv {
	return NewOapiSrv("isv", config)
}

// 企业内部
func NewCompanyOapiSrv(config *DTConfig) *OapiSrv {
	return NewOapiSrv("company", config)
}

// 普通钉钉用户帐号
func NewMiniOapiSrv(config *DTConfig) *OapiSrv {
	return NewOapiSrv("personalMini", config)
}
