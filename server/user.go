package server

import (
	"fmt"
	"net/url"

	pb "github.com/LingNovo/dingtalk/protos/oapi"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
)

// 通过Code获取userid
func (dtc *OapiSrv) UserIdByCode(ctx context.Context, req *pb.UserIdByCodeRequest) (*pb.UserIdByCodeResponse, error) {
	var data pb.UserIdByCodeResponse
	params := url.Values{}
	params.Add("code", req.Code) //requestAuthCode接口获取的CODE
	if err := dtc.httpRPC("user/getuserinfo", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 通过UnionId获取UserId
func (dtc *OapiSrv) UserIdByUnionId(ctx context.Context, req *pb.UserIdByUnionIdRequest) (*pb.UserIdByUnionIdResponse, error) {
	var data pb.UserIdByUnionIdResponse
	params := url.Values{}
	//用户在当前钉钉开发平台帐号范围内的唯一标识，
	//同一个钉钉开放平台帐号可以包含多个开放应用，
	//同时也包含ISV的套件应用及企业应用。
	params.Add("unionid", req.Unionid)
	if err := dtc.httpRPC("user/getUseridByUnionid", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 通过userid获取用户详细信息
// 通过CorpId+userId,可以在一个企业内唯一标记一个用户，
// 通过免登授权接口可以获取到当前用户的UserId信息
func (dtc *OapiSrv) UserInfoByUserId(ctx context.Context, req *pb.UserInfoByUserIdRequest) (*pb.UserInfoByUserIdResponse, error) {
	var data pb.UserInfoByUserIdResponse
	params := url.Values{}
	params.Add("lang", req.Lang)     //员工在企业内的UserId,企业用来唯一标识用户的字段。
	params.Add("userid", req.Userid) //通讯录语音（默认zh_CN,未来会支持en_US）
	if err := dtc.httpRPC("user/get", params, nil, &data, req.IsvCompanyInfo); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取部门成员（简化版）
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限;如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) UserSimpleList(ctx context.Context, req *pb.UserSimpleListRequest) (*pb.UserSimpleListResponse, error) {
	var data pb.UserSimpleListResponse
	params := url.Values{}
	if req.Lang != "" {
		params.Add("lang", req.Lang)
	}
	params.Add("department_id", fmt.Sprintf("%d", req.DepartmentId))
	if req.Size > 0 {
		// offset 从0开始
		params.Add("offset", fmt.Sprintf("%d", req.Offset))
		params.Add("size", fmt.Sprintf("%d", req.Size))
	}
	if req.Order != "" {
		params.Add("order", fmt.Sprintf("%s", req.Order))
	}
	if err := dtc.httpRPC("user/simplelist", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取部门成员（详情版）
// 如果你是ISV(应用服务商,将开发的应用上架到钉钉应用市场,提供给钉钉其他企业用户使用),则默认无调用权限。
func (dtc *OapiSrv) UserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	var data pb.UserListResponse
	params := url.Values{}
	if req.Lang != "" {
		params.Add("lang", req.Lang)
	}
	params.Add("department_id", fmt.Sprintf("%d", req.DepartmentId))
	if req.Size > 0 {
		// offset 从0开始
		params.Add("offset", fmt.Sprintf("%d", req.Offset))
		params.Add("size", fmt.Sprintf("%d", req.Size))
	}
	if req.Order != "" {
		params.Add("order", req.Order)
	}
	if err := dtc.httpRPC("user/list", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取管理员列表
func (dtc *OapiSrv) UserAdminList(ctx context.Context, req *google_protobuf.Empty) (*pb.UserAdminListResponse, error) {
	var data pb.UserAdminListResponse
	if err := dtc.httpRPC("user/get_admin", nil, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取管理员的微应用管理权限
// 如果你是ISV(应用服务商,将开发的应用上架到钉钉应用市场,提供给其他企业用户使用)
// 则只能查询自身的微应用权限
func (dtc *OapiSrv) UserCanAccessMicroapp(ctx context.Context, req *pb.UserCanAccessMicroappRequest) (*pb.UserCanAccessMicroappResponse, error) {
	var data pb.UserCanAccessMicroappResponse
	params := url.Values{}
	params.Add("appId", req.AppId)   //微应用id
	params.Add("userId", req.UserId) //员工唯一标识id
	if err := dtc.httpRPC("user/can_access_microap", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 创建成员
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其它企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉)，则有权限调用
// 本接口属于高权限接口，调用会被严格限制。请管理员在调用前完成个人实名认证，或者提交企业认证，人数上限将自动扩展
func (dtc *OapiSrv) UserCreate(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	var data pb.UserCreateResponse
	if err := dtc.httpRPC("user/create", nil, req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 更新成员
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) UserUpdate(ctx context.Context, req *pb.UserUpdateRequest) (*pb.OpenApiResponse, error) {
	var data pb.OpenApiResponse
	if err := dtc.httpRPC("user/update", nil, req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 删除成员
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) UserDelete(ctx context.Context, req *pb.UserDeleteRequest) (*pb.OpenApiResponse, error) {
	var data pb.OpenApiResponse
	params := url.Values{}
	params.Add("userid", req.Userid) //员工唯一标识ID（不可修改）
	if err := dtc.httpRPC("user/delete", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 批量删除成员
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) UserBatchDelete(ctx context.Context, req *pb.UserBatchDeleteRequest) (*pb.OpenApiResponse, error) {
	var data pb.OpenApiResponse
	body := map[string][]string{
		"useridlist": req.Useridlist, //员工UserID列表。列表长度在1到20之间
	}
	if err := dtc.httpRPC("user/batchdelete", nil, body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取企业员工人数
// 获取钉钉上某个企业/组织内的人数
func (dtc *OapiSrv) UserGetOrgUserCount(ctx context.Context, req *pb.UserGetOrgUserCountRequest) (*pb.UserGetOrgUserCountResponse, error) {
	var data pb.UserGetOrgUserCountResponse
	params := url.Values{}
	//0：包含未激活钉钉的人员数量 1：不包含未激活钉钉的人员数量
	params.Add("onlyActive", fmt.Sprintf("%d", req.OnlyActive))
	if err := dtc.httpRPC("user/get_org_user_count", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
