package service

import (
	"errors"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
	"xzdp-admin/system/lib/str"
)

type AdminService struct {
}

func (s *AdminService) Login(input dto.LoginReq) (bool, error) {
	//1.查询该用户是否存在
	db := myGrom.Db
	var adminUser entity.AdminUser
	if err := db.Model(&entity.AdminUser{}).Where("username", input.Username).Find(&adminUser).Error; err != nil {
		return false, err
	}
	if adminUser.ID == 0 {
		return false, errors.New("用户不存在")
	}
	//2.校验密码
	if !str.PasswordVerify(input.Password, adminUser.Password) {
		return false, errors.New("密码错误，请重试")
	}
	return true, nil
}

func (s *AdminService) Register(input dto.RegisterReq) error {
	db := myGrom.Db
	adminUser := entity.AdminUser{}
	//1.校验用户是否已存在
	if err := db.Model(&entity.AdminUser{}).Where("username", input.Username).Find(&adminUser).Error; err != nil {
		return err
	}
	if adminUser.ID != 0 {
		return errors.New("该用户已存在")
	}
	//2.创建用户
	adminUser.Username = input.Username
	adminUser.Name = input.Name
	adminUser.Phone = input.Phone
	adminUser.Remark = input.Remark
	adminUser.RoleID = input.RoleID
	adminUser.Password, _ = str.PasswordHash(input.Password)
	if err := db.Model(&entity.AdminUser{}).Create(&adminUser).Error; err != nil {
		return err
	}
	return nil
}
