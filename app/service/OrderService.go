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
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *OrderService) GetOrder(input dto.GetOrderInfoReq) (dto.OrderOutModel, error) {
	db := myGrom.Db
	var res dto.OrderOutModel
	if err := db.Model(entity.TbVoucherOrder{}).Where("id", input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
