package server

import (
	pb "github.com/LingNovo/dingtalk/protos/tapi"
	context "golang.org/x/net/context"
)

// 获取企业中角色下的员工列表
func (dtc *OapiSrv) RoleSimplelist(ctx context.Context, req *pb.RoleSimplelistRequest) (*pb.RoleSimplelistResponse, error) {
	general := TopMapRequest{
		"method":  corpRoleSimpleList,
		"role_id": req.RoleId, //角色ID
		"size":    req.Size,   //分页大小
		"offset":  req.Offset, //分页偏移
	}
	var data pb.RoleSimplelistResponse
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取企业角色列表
func (dtc *OapiSrv) RoleList(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListResponse, error) {
	var data pb.RoleListResponse
	general := TopMapRequest{
		"method": corpRoleList,
		"size":   req.Size,   //分页大小
		"offset": req.Offset, //分页偏移
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 企业在做企业员工管理的时候，需要对部分员工进行角色分类，该接口可以批量为员工增加角色信息
func (dtc *OapiSrv) AddRolesForEmps(ctx context.Context, req *pb.AddRolesForEmpsRequest) (*pb.AddRolesForEmpsResponse, error) {
	var data pb.AddRolesForEmpsResponse
	general := TopMapRequest{
		"method":       corpRoleAddRolesForemps,
		"rolelid_list": req.RolelidList, // 角色id list,最大长度为20
		"userid_list":  req.UseridList,  //员工id list，最大列表长度100
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 企业在做企业员工管理的时候，需要对部分员工进行角色分类，该接口可以批量删除员工的角色信息。 角色在钉钉的组织结构里面就是标签的意思，你可以批量为一批用户添加一批角色信息（dingtalk.corp.role.addrolesforemps），那么调用该接口就可以批量删除已经存在的角色和员工对应关系，角色和员工是多对多的关系。参考代码如下： req.setRolelidList("1,2,3,4,5"); // 已经存在的角色id列表 req.setUseridList("a,b,c,d,e"); // 用户的id列表
func (dtc *OapiSrv) RemoveRolesForEmps(ctx context.Context, req *pb.RemoveRolesForEmpsRequest) (*pb.RemoveRolesForEmpsResponse, error) {
	var data pb.RemoveRolesForEmpsResponse
	general := TopMapRequest{
		"method":      corpRoleRemoveRolesForemps,
		"roleid_list": req.RoleidList, //角色标签id，最大列表长度：20
		"userid_list": req.UseridList, //用户userId，最大列表长度：100
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 企业在做企业内部组织结构角色管理的时候，如果需要删除该企业下某个角色信息，可以调用该接口。调用的前提是该角色下面的所有员工都已经被删除该角色
func (dtc *OapiSrv) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	var data pb.DeleteRoleResponse
	general := TopMapRequest{
		"method":  corpRoleDeleteRole,
		"role_id": req.RoleId, // 角色Id
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 该接口通过groupId参数可以获取该角色组详细信息以及下面所有关联的角色的信息
func (dtc *OapiSrv) GetRoleGroup(ctx context.Context, req *pb.GetRoleGroupRequest) (*pb.GetRoleGroupResponse, error) {
	var data pb.GetRoleGroupResponse
	general := TopMapRequest{
		"method":   corpRoleGetRoleGroup,
		"group_id": req.GroupId, // 角色组的Id
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取角色详情
func (dtc *OapiSrv) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	var data pb.GetRoleResponse
	general := TopMapRequest{
		"method": corpRoleGetRole,
		"roleId": req.RoleId,
	}
	if err := dtc.httpTOP(general, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
