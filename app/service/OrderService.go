package service

import (
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
)

type OrderService struct {
}

func (s *OrderService) OrderList(input dto.OrderListReq) (dto.OrderListRes, error) {
	db := myGrom.Db
	var res dto.OrderListRes
	query := entity.TbVoucherOrder{
		UserID:  input.UserID,
		PayType: input.PayType,
		Status:  input.Status,
	}
	db = db.Model(&entity.TbVoucherOrder{}).Where(&query)
	db = db.Joins("left join tb_user user on user_id = user.id")
	db = db.Joins("left join tb_voucher voucher on voucher_id = voucher.id")
	db = db.Select("tb_voucher_order.*, user.nick_name, voucher.title")
	if input.UserName != "" {
		db = db.Where("user.nick_name like ?", "%"+input.UserName+"%")
	}
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *OrderService) GetOrder(input dto.GetOrderInfoReq) (dto.OrderOutModel, error) {
	db := myGrom.Db
	var res dto.OrderOutModel
	db = db.Model(entity.TbVoucherOrder{})
	db = db.Joins("left join tb_user user on user_id = user.id")
	db = db.Joins("left join tb_voucher voucher on voucher_id = voucher.id")
	db = db.Select("tb_voucher_order.*, user.nick_name, voucher.title")
	if err := db.Where("tb_voucher_order.id", input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
