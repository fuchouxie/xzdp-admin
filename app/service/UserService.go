package service

import (
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
)

type UserService struct {
}

func (s *UserService) UserList(input dto.UserListReq) (dto.UserListRes, error) {
	db := myGrom.Db
	var res dto.UserListRes
	db = db.Model(&entity.TbUser{}).Joins("left join tb_user_info userInfo on tb_user.id = userInfo.user_id")
	db = db.Select("tb_user.*, tb_user.created_at, userInfo.city, userInfo.level")
	if input.NickName != "" {
		db = db.Where("nick_name like ?", "%"+input.NickName+"%")
	}
	if input.Phone != "" {
		db = db.Where("phone like ?", "%"+input.Phone+"%")
	}
	if input.Level != -1 {
		db = db.Where("userInfo.level", input.Level)
	}
	if input.City != "" {
		db = db.Where("userInfo.city like ?", "%"+input.City+"%")
	}
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *UserService) Remove(input dto.RemoveUserReq) error {
	db := myGrom.Db
	if err := db.Where("id", input.ID).Delete(&entity.TbUser{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(input dto.GetUserInfoReq) (dto.UserOutModel, error) {
	db := myGrom.Db
	var res dto.UserOutModel
	db = db.Model(&entity.TbUser{}).Joins("left join tb_user_info userInfo on tb_user.id = userInfo.user_id")
	db = db.Select("tb_user.*, userInfo.city, userInfo.level, userInfo.introduce, userInfo.fans, userInfo.followee" +
		", userInfo.gender, userInfo.birthday, userInfo.credits")
	if err := db.Where("tb_user.id", input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
