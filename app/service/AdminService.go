package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
	"xzdp-admin/system/core/mySession"
	"xzdp-admin/system/lib/str"
)

type AdminService struct {
}

func (s *AdminService) Code(ctx *gin.Context, input dto.CodeReq) (string, error) {
	db := myGrom.Db
	var adminUser entity.AdminUser
	if err := db.Model(&entity.AdminUser{}).Where("phone", input.Phone).Find(&adminUser).Error; err != nil {
		return "手机号有误请重试", err
	}
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	adminUser.Code = code
	if err := db.Model(&entity.AdminUser{}).Where("id", adminUser.ID).Updates(&adminUser).Error; err != nil {
		return "", err
	}
	return code, nil
}

func (s *AdminService) CodeLogin(ctx *gin.Context, input dto.CodeLoginReq) error {
	db := myGrom.Db
	var adminUser entity.AdminUser
	if err := db.Model(&entity.AdminUser{}).Where("phone", input.Phone).Find(&adminUser).Error; err != nil {
		return err
	}
	if input.Code != adminUser.Code {
		return errors.New("验证码有误")
	}
	mySession.SaveSession(ctx, "userId", fmt.Sprintf("%d", adminUser.ID))
	return nil
}

func (s *AdminService) Login(ctx *gin.Context, input dto.LoginReq) (bool, error) {
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
	//3.保存用户id到session
	mySession.SaveSession(ctx, "userId", fmt.Sprintf("%d", adminUser.ID))
	return true, nil
}

func (s *AdminService) Register(input dto.RegisterReq) error {
	db := myGrom.Db
	adminUser := entity.AdminUser{}
	//1.校验用户是否已存在
	if err := db.Model(&entity.AdminUser{}).Unscoped().Where("username", input.Username).Find(&adminUser).Error; err != nil {
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

func (s *AdminService) UpdateInfo(input dto.UpdateAdminReq) error {
	db := myGrom.Db
	adminUser := entity.AdminUser{
		ID:       input.ID,
		Username: input.Username,
		Phone:    input.Phone,
		RoleID:   input.RoleID,
		Name:     input.Name,
		Remark:   input.Remark,
	}
	if err := db.Model(&entity.AdminUser{}).Where("id", input.ID).Updates(&adminUser).Error; err != nil {
		return err
	}
	return nil
}

func (s *AdminService) AdminList(input dto.AdminListReq) (dto.AdminListRes, error) {
	db := myGrom.Db
	var res dto.AdminListRes
	query := entity.AdminUser{
		ID:     input.ID,
		RoleID: input.RoleID,
	}
	db = db.Model(&entity.AdminUser{}).Where(query)
	if input.Username != "" {
		db.Where("username like ?", "%"+input.Username+"%")
	}
	if input.Phone != "" {
		db.Where("phone like ?", "%"+input.Phone+"%")
	}
	if input.Name != "" {
		db.Where("name like ?", "%"+input.Name+"%")
	}
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

//func (s *AdminService) ChangePassword(input dto.ChangePasswordReq) error {
//	db := myGrom.Db
//	//1.验证用户是否存在
//	adminUser := entity.AdminUser{}
//	if err := db.Model(&entity.AdminUser{}).Where("username", input.Username).Find(&adminUser).Error; err != nil {
//		return err
//	}
//	if adminUser.ID == 0 {
//		return errors.New("用户不存在")
//	}
//	//2.校验旧密码是否正确
//	if !str.PasswordVerify(input.OldPassword, adminUser.Password) {
//		return errors.New("密码有误")
//	}
//	//3.校验新密码是否与旧密码一致
//	if input.NewPassword == input.OldPassword {
//		return errors.New("新密码不能与旧密码一致")
//	}
//	//加密
//	adminUser.Password, _ = str.PasswordHash(input.NewPassword)
//	//4.执行修改
//	if err := db.Model(&entity.AdminUser{}).Where("id", adminUser.ID).Update("password", adminUser.Password).Error; err != nil {
//		return err
//	}
//	return nil
//}

func (s *AdminService) ChangePasswordV2(userInfo *entity.AdminUser, input dto.ChangePasswordReq) error {
	db := myGrom.Db
	//2.校验旧密码是否正确
	if !str.PasswordVerify(input.OldPassword, userInfo.Password) {
		return errors.New("密码有误")
	}
	//3.校验新密码是否与旧密码一致
	if input.NewPassword == input.OldPassword {
		return errors.New("新密码不能与旧密码一致")
	}
	//加密
	userInfo.Password, _ = str.PasswordHash(input.NewPassword)
	//4.执行修改
	if err := db.Model(&entity.AdminUser{}).Where("id", userInfo.ID).Update("password", userInfo.Password).Error; err != nil {
		return err
	}
	return nil
}

func (s *AdminService) Remove(input dto.RemoveAdminReq) error {
	db := myGrom.Db
	if err := db.Model(&entity.AdminUser{}).Delete("id", input.ID).Error; err != nil {
		return err
	}
	return nil
}

func (s *AdminService) BatchRemove(input dto.BatchRemoveAdminReq) error {
	db := myGrom.Db
	ids := str.IntStrToArray(input.IDS, ",")
	if err := db.Where("id", ids).Delete(&entity.AdminUser{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *AdminService) GetAdminUser(input dto.GetAdminInfoReq) (entity.AdminUser, error) {
	db := myGrom.Db
	var res entity.AdminUser
	if err := db.Model(&entity.AdminUser{}).Where("id", input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
