package tapi

import (
	"fmt"
)

func (m *TopApiResponse) CheckError() error {
	if m.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.Code, m.Msg, m.SubCode, m.SubMsg)
	}
	return nil
}

func (m *RoleSimplelistResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *RoleListResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *AddRolesForEmpsResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *RemoveRolesForEmpsResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *DeleteRoleResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *GetRoleGroupResponse) CheckError() error {
	if m.ErrorResponse == nil {
		return nil
	}
	if m.ErrorResponse.Code > 0 {
		return fmt.Errorf("code: %d\nmsg: %s\nsub_code: %s\nsub_msg: %s", m.ErrorResponse.Code, m.ErrorResponse.Msg, m.ErrorResponse.SubCode, m.ErrorResponse.SubMsg)
	}
	return nil
}

func (m *GetRoleResponse) CheckError() error {
	if m.DingtalkOapiRoleGetroleResponse == nil {
		return nil
	}
	if m.DingtalkOapiRoleGetroleResponse.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.DingtalkOapiRoleGetroleResponse.Errcode, m.DingtalkOapiRoleGetroleResponse.Errmsg)
	}
	return nil
}
