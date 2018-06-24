package server

import (
	"fmt"
	"net/url"

	pb "github.com/LingNovo/dingtalk/protos/oapi"
	context "golang.org/x/net/context"
)

// 获取子部门Id列表
// 注意：该接口能拉取到企业部门下的所有直属子部门列表。不受授权范围限制，ISV可以根据该接口完成企业部门的遍历。
// ISV默认无调用权限
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) SubDepartmentList(ctx context.Context, req *pb.SubDepartmentListRequest) (*pb.SubDepartmentListResponse, error) {
	var data pb.SubDepartmentListResponse
	params := url.Values{}
	params.Add("id", fmt.Sprint(req.Id)) //父部门id(如果不传，默认部门为根部门，根部门ID为1)
	if err := dtc.httpRPC("department/list_ids", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取部门列表
// 注意：企业开发者所使用的CorpSecret都是按照部门进行设置的，企业开发者需要先查询自己的CorpSecret对应的授权部门列表
// ISV默认无调用权限
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) DepartmentList(ctx context.Context, req *pb.DepartmentListRequest) (*pb.DepartmentListResponse, error) {
	var data pb.DepartmentListResponse
	params := url.Values{}
	if req.Lang != "" {
		params.Add("lang", req.Lang)
	}
	if req.FetchChild {
		params.Add("fetch_child", fmt.Sprintf("%t", req.FetchChild))
	}
	if req.Id != 0 {
		params.Add("id", fmt.Sprint(req.Id))
	}
	if err := dtc.httpRPC("department/list", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取部门详情
// ISV默认无调用权限
// 如果你是ISV(应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用)，则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），有权限调用
func (dtc *OapiSrv) DepartmentDetail(ctx context.Context, req *pb.DepartmentDetailRequest) (*pb.DepartmentDetailResponse, error) {
	var data pb.DepartmentDetailResponse
	params := url.Values{}
	params.Add("id", fmt.Sprint(req.Id)) //部门id
	params.Add("lang", req.Lang)         //通讯录语言(默认zh_CN，未来会支持en_US)
	if err := dtc.httpRPC("department/get", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 创建部门
// ISV默认无调用权限
// 如果你是ISV（应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用），则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），有限调用
func (dtc *OapiSrv) DepartmentCreate(ctx context.Context, req *pb.DepartmentCreateRequest) (*pb.DepartmentCreateResponse, error) {
	var data pb.DepartmentCreateResponse
	if err := dtc.httpRPC("department/create", nil, req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 更新部门
// 如果你是ISV（应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用），则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），则有权限调用
func (dtc *OapiSrv) DepartmentUpdate(ctx context.Context, req *pb.DepartmentUpdateRequest) (*pb.DepartmentUpdateResponse, error) {
	var data pb.DepartmentUpdateResponse
	if err := dtc.httpRPC("department/update", nil, req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 删除部门
// 如果你是ISV（应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用），则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），有权限调用
func (dtc *OapiSrv) DepartmentDelete(ctx context.Context, req *pb.DepartmentDeleteRequest) (*pb.OpenApiResponse, error) {
	var data pb.OpenApiResponse
	params := url.Values{}
	//部门id。（注：不能删除根部门；不能删除含有子部门、成员的部门）
	params.Add("id", fmt.Sprintf("%d", req.Id))
	if err := dtc.httpRPC("department/delete", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil

}

// 查询部门的所有上级父部门路径
// 如果你是ISV（应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用），则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），有权限调用
/*
假设部门的组织结构如下
1
  |->123
    |->456
      |->789
当传入部门id为789时，返回的结果按顺序依次为当前部门id及其所有父部门的ID，直到根部门，如[789,456,123,1]。
*/
func (dtc *OapiSrv) DepartmentListParentDeptsByDept(ctx context.Context, req *pb.DepartmentListParentDeptsByDeptRequest) (*pb.DepartmentListParentDeptsByDeptResponse, error) {
	var data pb.DepartmentListParentDeptsByDeptResponse
	params := url.Values{}
	//希望查询的部门的id，包含查询的部门本身
	params.Add("id", fmt.Sprintf("%d", req.Id))
	if err := dtc.httpRPC("department/list_parent_depts_by_dept", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 查询指定用户的所有上级父部门路径
// 如果你是ISV（应用服务商，将开发的应用上架到钉钉应用市场，提供给钉钉其他企业用户使用），则默认无调用权限
// 如果你是企业内部开发者（将自己公司的HR、OA、客户管理、业务管理等系统接入钉钉），有权限调用
/*
假设用户所属部门的组织结构如下
1
  |->123
    |->456  ->员工A
  |->789  ->员工A
当传入员工A的userId时，返回的结果按顺序依次为其所有父部门的ID，直到根部门，如[[456,123,1],[789,1]]。
*/
func (dtc *OapiSrv) DepartmentListParentDepts(ctx context.Context, req *pb.DepartmentListParentDeptsRequest) (*pb.DepartmentListParentDeptsResponse, error) {
	var data pb.DepartmentListParentDeptsResponse
	params := url.Values{}
	params.Add("userid", req.UserId) //希望查询的用户的id
	if err := dtc.httpRPC("department/list_parent_depts", params, nil, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
