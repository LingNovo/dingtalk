package server

import (
	"fmt"
	"net/url"
	"time"

	pb "github.com/LingNovo/dingtalk/protos/oapi"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
)

type AccessTokenResponse struct {
	pb.OpenApiResponse
	AccessToken string `json:"access_token"` //获取到的凭证
	Expires     int    `json:"expires_in"`   //过期时间
	Created     int64
}

type SSOAccessTokenResponse struct {
	pb.OpenApiResponse
	SSOAccessToken string `json:"access_token"`
	Expires        int    `json:"expires_in"`
	Created        int64
}

type SNSAccessTokenResponse struct {
	pb.OpenApiResponse
	SNSAccessToken string `json:"access_token"`
	Expires        int    `json:"expires_in"`
	Created        int64
}

type SuiteAccessTokenResponse struct {
	pb.OpenApiResponse
	SuiteAccessToken string `json:"suite_access_token"` //应用套件access_token
	Expires          int    `json:"expires_in"`         //有效期7200秒，过期之前要主动更新（建议ISV服务端做定时器主动更新，而不是依赖钉钉的定时推送）
	Created          int64
}

type TicketResponse struct {
	pb.OpenApiResponse
	Ticket  string `json:"ticket"`     //用于JS API的临时票据
	Expires int    `json:"expires_in"` //票据过期时间
	Created int64
}

type GetPermanentCodeResponse struct {
	pb.OpenApiResponse
	PermanentCode   string       `json:"permanent_code"`
	ChPermanentCode string       `json:"ch_permanent_code"`
	AuthCorpInfo    AuthCorpInfo `json:"auth_corp_info"`
}

type AuthCorpInfo struct {
	CorpId   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

type ActivateSuiteResponse struct {
	pb.OpenApiResponse
}

type GetCorpAccessTokenResponse struct {
	pb.OpenApiResponse
	//授权方（企业）corp_access_token
	AccessToken string `json:"access_token"`
	//授权方（企业）corp_access_token超时时间
	Expires int `json:"expires_in"`
}

type GetCompanyInfoResponse struct {
	pb.OpenApiResponse
	//授权方企业信息
	AuthCorpInfo GCIRAuthCorpInfo `json:"auth_corp_info"`
	//授权方管理员信息
	AuthUserInfo GCIRAuthUserInfo `json:"auth_user_info"`
	//授权的应用信息
	AuthInfo GCIRAuthInfo `json:"auth_info"`
	//授权的服务窗应用信息列表
	ChannelAuthInfo GCIRChannelAuthInfo `json:"channel_auth_info"`
}

type GCIRAuthCorpInfo struct {
	//授权方企业id
	CorpId string `json:"corpid"`
	//邀请码，只有自己邀请的企业才会返回邀请码，可用该邀请码统计不同渠道拉新，否则值为空字符串
	InviteCode string `json:"invite_code"`
	//企业所属行业
	Industry string `json:"industry"`
	//授权方企业名称
	CorpName string `json:"corp_name"`
	//序列号
	LicenseCode string `json:"license_code"`
	//渠道码
	AuthChannel string `json:"auth_channel"`
	//渠道类型,为了避免渠道码重复，可与渠道码共同确认渠道（可能为空。非空时当前只有满天星类型，值为STAR_ACTIVITY）
	AuthChannelType string `json:"auth_channel_type"`
	//企业是否认证
	IsAuthenticated bool `json:"is_authenticated"`
	//企业认证等级，0：未认证，1：高级认证，2：中级认证，3：初级认证
	AuthLevel int `json:"auth_level"`
	//企业邀请链接
	InviteURL string `json:"invite_url"`
	//企业logo
	CorpLogoURL string `json:"corp_logo_url"`
}

type GCIRAuthUserInfo struct {
	//员工唯一标识Id（不可修改）
	UserID string `json:"userId"`
}

// 授权的应用信息
type GCIRAuthInfo struct {
	Agent []GCIRAuthInfoAgent `json:"agent"`
}

type GCIRAuthInfoAgent struct {
	//授权方应用名字
	AgentName string `json:"agent_name"`
	//授权方应用id ，agentId会在前端签名校验的时候使用到。
	AgentId int `json:"agentid"`
	//套件中的应用id
	AppId int `json:"appid"`
	//授权方应用头像
	LogoUrl string `json:"logo_url"`
	//对此微应用有管理权限的管理员userid
	AdminList string `json:"admin_list"`
}

// 授权的服务窗应用信息列表
type GCIRChannelAuthInfo struct {
	ChannelAgent []GCIRChannelAuthInfoAgent `json:"channelAgent"`
}

type GCIRChannelAuthInfoAgent struct {
	//授权方应用名字
	AgentName string `json:"agent_name"`
	//授权方应用id ，agentId会在前端签名校验的时候使用到。
	AgentId string `json:"agentid"`
	//套件中的应用id
	AppId string `json:"appid"`
	//对此微应用有管理权限的管理员userid
	LogoUrl string `json:"logo_url"`
}

type ScopesResponse struct {
	pb.OpenApiResponse
	//可以得到的企业用户字段
	//name:员工名称,email员工邮箱 具体参见user/get接口字段描述
	AuthUserField []string `json:"auth_user_field"`
	//ISV可以直接使用企业的功能字段
	//contact_call:isv是否可以电话联系企业管理员
	ConditionField []string      `json:"condition_field"`
	AuthOrgScopes  AuthOrgScopes `json:"auth_org_scopes"`
}

type AuthOrgScopes struct {
	//企业授权的部门id列表
	//返回值为授权部门id的并集。 (设置“全部员工”时，返回授权的部门id为根部门ID,员工userid列表为空)
	AuthedDept []int `json:"authed_dept"`
	//企业授权的员工userid列表
	//返回值为授权人员id的并集。(设置“仅为管理员可见”时，返回所有的管理员id。授权的部门id列表为空)
	AuthedUser []string `json:"authed_user"`
}

func (e *AccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *AccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *TicketResponse) CreatedAt() int64 {
	return e.Created
}

func (e *TicketResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SSOAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SSOAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SNSAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SNSAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SuiteAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SuiteAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

// 刷新企业获取的access_token
// Access_Token是企业访问钉钉开放平台全局接口的唯一凭证，即调用接口时需携带Access_Token。
// 对于企业接入来说，AccessToken需要用CorpID和CorpSecret来换取，CorpID是企业在钉钉中的标识；每个企业拥有一个唯一的CorpID，CorpSecret是企业每个微应用的凭证密钥。
// CorpID和CorpSecret可以在开发者后台中找到，
// 对于ISV接入来说，可以通过企业对ISV的授权获取企业授权的access_token。
// 不同的CorpSecret会返回不同的AccessToken。正常情况下AccessToken有效期为7200秒，有效期内重复获取返回相同结果，并自动续期。
func (dtc *OapiSrv) RefreshCompanyAccessToken(ctx context.Context, req *google_protobuf.Empty) (*google_protobuf.Empty, error) {
	dtc.Locker.Lock()
	var data AccessTokenResponse
	err := dtc.AccessTokenCache.Get(&data)
	if err == nil {
		dtc.AccessToken = data.AccessToken
		fmt.Printf("Get access_token To Local Cache=%s\n", dtc.AccessToken)
		dtc.Locker.Unlock()
		return &google_protobuf.Empty{}, nil
	}
	params := url.Values{}
	//企业Id
	params.Add("corpid", dtc.DTConfig.CorpId)
	//企业应用的凭证密钥
	params.Add("corpsecret", dtc.DTConfig.CorpSecret)
	err = dtc.httpRPC("gettoken", params, nil, &data)
	if err == nil {
		dtc.AccessToken = data.AccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.AccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return &google_protobuf.Empty{}, err
}

// 刷新企业获取的sso_access_token
// 开发者在开发微应用的时候，除了需要开发移动端、PC端应用以外，同样需要开发微应用的管理后台。
// 获取免登SsoToken
// 本接口获取的SsoToken和上面获取的AccessToken应用场景不一样，SsoToken只在微应用后台管理免登服务中使用。
// 注意：ISV开发的微应用后台免登中的SsoToken，使用ISV自己的corpid和ssosecret来换取，不是使用用户企业的。
func (dtc *OapiSrv) RefreshSSOAccessToken() error {
	dtc.Locker.Lock()
	var data SSOAccessTokenResponse
	err := dtc.SSOAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SSOAccessToken = data.SSOAccessToken
		fmt.Printf("Get sso_access_token To Local Cache=%s\n", dtc.SSOAccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	params := url.Values{}
	//企业Id
	params.Add("corpid", dtc.DTConfig.CorpId)
	//这里必须填写专属的SSOSecret
	params.Add("corpsecret", dtc.DTConfig.SSOSecret)
	err = dtc.httpSSO("sso/gettoken", params, nil, &data)
	if err == nil {
		dtc.SSOAccessToken = data.SSOAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SSOAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 刷新 SNS access_token
// 普通钉钉用户账号开放
// 获取钉钉开放应用的ACCESS_TOKEN
// 普通钉钉用户账号开放相关接口和企业应用开发、ISV应用开发无关，第三方web服务提供商取得钉钉开放应用的appid及appsecret后，可以获取开放应用的ACCESS_TOKEN
func (dtc *OapiSrv) RefreshSNSAccessToken(ctx context.Context, req *google_protobuf.Empty) (*google_protobuf.Empty, error) {
	dtc.Locker.Lock()
	var data SNSAccessTokenResponse
	err := dtc.SNSAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SNSAccessToken = data.SNSAccessToken
		fmt.Printf("Get sns_access_token To Local Cache=%s\n", dtc.SNSAccessToken)
		dtc.Locker.Unlock()
		return &google_protobuf.Empty{}, nil
	}
	params := url.Values{}
	//由钉钉开放平台提供给开放应用的唯一标识
	params.Add("appid", dtc.DTConfig.SNSAppId)
	//由钉钉开放平台提供的密钥
	params.Add("appsecret", dtc.DTConfig.SNSSecret)
	err = dtc.httpSNS("sns/gettoken", params, nil, &data)
	if err == nil {
		dtc.SNSAccessToken = data.SNSAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SNSAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return &google_protobuf.Empty{}, err
}

// 刷新 isv suite_access_token
// 获取应用套件令牌Token
/*
该API用于获取应用套件令牌（suite_access_token）
1.每次获取suite_access_token的时候，都会拿到suite_access_token的过期时间。在使用这个suite_access_token的时候，需要先判断一下suite_access_token是否过期(例如差10分钟过期，或者已经过期了)，如果发现suite_access_token即将过期需要再来获取一下suite_access_token，获取suite_access_token的时候要加锁（防止多个请求同时都在更新suite_access_token）。
2.每次获取corp_access_token的时候，都会拿到corp_access_token的过期时间。在使用这个corp_access_token的时候，需要先判断一下corp_access_token是否过期(例如差10分钟过期或者已经过期了)，如果发现corp_access_token即将过期需要再来获取一下corp_access_token，获取corp_access_token的时候要加锁。
3.有些ISV更新suite_access_token是依赖suiteticket推送来更新的，不建议这样来实现。钉钉的suite_access_token是不定时推送的，依赖suite_access_token的推送会引起suitetoken失效，导致程序出现异常。
*/
func (dtc *OapiSrv) RefreshSuiteAccessToken() error {
	dtc.Locker.Lock()
	var data SuiteAccessTokenResponse
	err := dtc.SuiteAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SuiteAccessToken = data.SuiteAccessToken
		fmt.Printf("Get suite_access_token To Local Cache=%s\n", dtc.SuiteAccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	info := map[string]string{
		//套件key，开发者后台创建套件后获取
		"suite_key": dtc.DTConfig.SuiteKey,
		//套件secret，开发者后台创建套件后获取
		"suite_secret": dtc.DTConfig.SuiteSecret,
		//钉钉推送的ticket
		"suite_ticket": dtc.DTConfig.SuiteTicket,
	}
	err = dtc.httpSNS("service/get_suite_token", nil, info, &data)
	if err == nil {
		dtc.SuiteAccessToken = data.SuiteAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SuiteAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 获取Ticket
//企业在使用微应用中的JS API时，需要先从钉钉开放平台接口获取jsapi_ticket生成签名数据，并将最终签名用的部分字段及签名结果返回到H5中，JS API底层将通过这些数据判断H5是否有权限使用JS API。
/*
注意：
1.在企业应用中，在jsticket未过期的时候通过get_jsapi_ticket获取到的都是一个全新的jsticket（和旧的jsticket值不同），这个全新的jsticket过期时间是2小时。
2.在ISV应用中，在jsticket未过期的时候，也就是两小时之内通过get_jsapi_ticket获取到的jsticket和老的jsticket值相同，只是过期时间延长到2小时。
3.jsticket是以一个企业对应一个，所以在使用的时候需要将jsticket以企业为维度进行缓存下来（设置缓存过期时间2小时），并不需要每次都通过接口拉取。
*/
func (dtc *OapiSrv) GetJSAPITicket() (ticket string, err error) {
	dtc.Locker.Lock()
	var data TicketResponse
	err = dtc.TicketCache.Get(&data)
	if err == nil {
		dtc.Locker.Unlock()
		return data.Ticket, err
	}
	err = dtc.httpRPC("get_jsapi_ticket", nil, nil, &data)
	if err == nil {
		ticket = data.Ticket
		dtc.TicketCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return ticket, err
}

// 获取钉钉配置信息（JSAPI）
func (dtc *OapiSrv) DingTalkConfig(ctx context.Context, req *pb.DingTalkConfigRequest) (*pb.DingTalkConfigResponse, error) {
	var data pb.DingTalkConfigResponse
	ticket, err := dtc.GetJSAPITicket()
	if err != nil {
		return nil, err
	}
	data.TargetUrl = req.TargetUrl
	data.NonceStr = randomString(32)
	data.AgentId = dtc.DTConfig.AgentId
	data.TimeStamp = fmt.Sprint(time.Now().Unix())
	data.CorpId = dtc.DTConfig.CorpId
	data.Ticket = ticket
	data.Signature = sign(ticket, data.NonceStr, data.TimeStamp, req.TargetUrl)
	return &data, nil
}

// 签名
func sign(ticket string, nonceStr string, timeStamp string, url string) string {
	s := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonceStr, timeStamp, url)
	return sha1Sign(s)
}

// 获取CorpSecret授权范围
/*
钉钉已经支持企业管理员动态调整ISV开发者用到的token及可获取到的企业通讯录范围(之前token是可以获取企业全部通讯录组织架构的)，在ISV开发者获取到可以访问企业通讯录接口的token后，在正式获取通讯录数据前，首先应该调用本接口，获取当前token能够读取到的通讯录部门及人员范围，再依次遍历本接口返回的部门id列表，调用通讯录接口获取每个部门及其子部门的数据。另外，在套件收到【向ISV推送授权变更消息】回调后，也需要调用本接口来同步通讯录数据。
举个例子，某企业的通讯录包含两个部门和三个员工，分别是部门A(员工甲、员工乙)和部门B(员工丙)。
以前，在企业对ISV的套件授权后，ISV可以通过/department/list接口获取部门A和部门B，再通过/user/simplelist接口获取到员工甲、员工乙及员工丙的userid，最后通过/user/get获取所有三个员工的详细信息，完成通讯录同步。
现在，在企业对ISV的套件授权时，管理员可以选择只授权一部分通讯录数据给ISV，针对本例，假如只授权了员工甲及部门B，当ISV再调用/department/list接口时，部门A和部门B的id都不再返回，会返回错误[部门不在授权范围内]。这会导致ISV无法获取管理员授予的全部数据，这时ISV应该调用本次新增的/auth/scopes接口获取管理员授予的全部数据范围，接口会返回部门B的id及员工甲的userid，然后isv再通过/department/list?id={部门B的id}获取部门B的子部门(本例中为空)，通过/user/get接口获取员工甲的详细信息。
需要注意的是，如果ISV在通过企业会话发送消息时，如果接收者是未授权的员工userid或部门id，也会返回错误[员工或部门不在授权范围内]。
为了方便ISV平滑迁移，我们暂定在2017年4月10日之后新创建的套件默认采用按部门授权的策略，之前创建的套件，如需要开启，请通过工单形式提供suiteKey，由我们手动开启。需要注意的是，在正式套件开启前，一定要先建立一个测试套件，对功能完整测试后，再申请对正式套件开启。
在套件开启按部门授权功能后，企业管理员可以在如下位置进行授权范围的调整，钉钉客户端->工作TAB->管理->选择应用->设置->授权使用范围。
*/
func (dtc *OapiSrv) GetAuthScopes() (ScopesResponse, error) {
	var data ScopesResponse
	err := dtc.httpRPC("auth/scopes", nil, nil, &data)
	return data, err
}

// 获取企业授权的永久授权码
/*
该API用于使用临时授权码换取企业的永久授权码，并可通过该永久授权码换取该企业授权信息、企业access_token。请将永久授权码存储下来，以后该企业的corp_access_token每次更新都需要使用。
注：临时授权码使用一次后即失效
*/
func (dtc *OapiSrv) IsvGetPermanentCode(tmpAuthCode string) (GetPermanentCodeResponse, error) {
	var data GetPermanentCodeResponse
	requestData := map[string]string{
		//回调接口（tmp_auth_code）获取的临时授权码
		"tmp_auth_code": tmpAuthCode,
	}
	err := dtc.httpRPC("service/get_permanent_code", nil, requestData, &data)
	return data, err
}

// 激活套件
/*
该API适用于当企业用户授权开通套件时，为授权方企业激活套件。如果ISV未调用此接口，则企业用户无法使用ISV套件。
注意：为企业激活该套件之后，就可以在钉钉客户端的工作面板上面看到该应用。
*/
func (dtc *OapiSrv) IsvActivateSuite(authCorpID string, permanentCode string) (ActivateSuiteResponse, error) {
	var data ActivateSuiteResponse
	requestData := map[string]string{
		//套件key
		"suite_key": dtc.DTConfig.SuiteKey,
		//授权方corpid
		"auth_corpid": authCorpID,
		//永久授权码，从get_permanent_code接口中获取
		"permanent_code": permanentCode,
	}
	err := dtc.httpRPC("service/activate_suite", nil, requestData, &data)
	return data, err
}

// 获取企业授权的凭证
// 企业的授权凭证access_token可以用来调用企业相关的数据接口。
func (dtc *OapiSrv) IsvGetCorpAccessToken(authCorpID string, permanentCode string) (GetCorpAccessTokenResponse, error) {
	var data GetCorpAccessTokenResponse
	requestData := map[string]string{
		//授权方corpid
		"auth_corpid": authCorpID,
		//永久授权码，通过get_permanent_code获取
		"permanent_code": permanentCode,
	}
	err := dtc.httpRPC("service/get_corp_token", nil, requestData, &data)
	return data, err
}

// 直接获取企业授权的凭证
func (dtc *OapiSrv) IsvGetCAT(tmpAuthCode string) {

}

// 获取企业的基本信息
// 该API用于通过suite_key和企业的corpId获取企业的基本信息
func (dtc *OapiSrv) IsvGetCompanyInfo(authCorpID string) (GetCompanyInfoResponse, error) {
	var data GetCompanyInfoResponse
	requestData := map[string]string{
		//授权方corpid
		"auth_corpid": authCorpID,
		//套件key
		"suite_key": dtc.DTConfig.SuiteKey,
	}
	err := dtc.httpRPC("service/get_auth_info", nil, requestData, &data)
	return data, err
}
