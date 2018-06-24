package oapi

import (
	"fmt"
)

func (m *OpenApiResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserIdByCodeResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserIdByUnionIdResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserInfoByUserIdResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserSimpleListResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserListResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserAdminListResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserCanAccessMicroappResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserCreateResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *UserGetOrgUserCountResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *SubDepartmentListResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentListResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentDetailResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentCreateResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentUpdateResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentListParentDeptsByDeptResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *DepartmentListParentDeptsResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *SnsGetPersistentCodeResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *SnsGetSnsTokenResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}

func (m *SnsGetUserInfoResponse) CheckError() error {
	if m.Errcode != 0 {
		return fmt.Errorf("errcode: %d\nerrmsg: %s", m.Errcode, m.Errmsg)
	}
	return nil
}
