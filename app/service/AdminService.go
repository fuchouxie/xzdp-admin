package service

import "xzdp-admin/app/dto"

type AdminService struct {
}

func (s AdminService) Login(input dto.LoginReq) (bool, error) {
	return true, nil
}
