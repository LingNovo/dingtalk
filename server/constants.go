package server

// 包外部可用的常量
const (
	VERSION   = "0.1"
	OAPIURL   = "https://oapi.dingtalk.com/"
	TOPAPIURL = "https://eco.taobao.com/router/rest"
)

// 包内部用的常量
const (
	signMD5         = "MD5"
	signHMAC        = "HMAC"
	topFormat       = "json"
	topVersion      = "2.0"
	topSimplify     = false
	topSecret       = "YukNetwork"
	topSignMethod   = signMD5
	typeJSON        = "application/json"
	typeForm        = "application/x-www-form-urlencoded"
	typeMultipart   = "multipart/form-data"
	aesEncodeKeyLen = 43
)

// Top接口
const (
	corpRoleSimpleList         = "dingtalk.corp.role.simplelist"
	corpRoleList               = "dingtalk.corp.role.list"
	corpRoleAddRolesForemps    = "dingtalk.corp.role.addrolesforemps"
	corpRoleRemoveRolesForemps = "dingtalk.corp.role.removerolesforemps"
	corpRoleDeleteRole         = "dingtalk.corp.role.deleterole"
	corpRoleGetRoleGroup       = "dingtalk.corp.role.getrolegroup"
	corpRoleGetRole            = "dingtalk.oapi.role.getrole"
)
