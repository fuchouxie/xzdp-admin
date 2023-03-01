package service

import (
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
	"xzdp-admin/system/lib/str"
)

type VoucherService struct {
}

func (s *VoucherService) VoucherList(input dto.VoucherListReq) (dto.VoucherListRes, error) {
	db := myGrom.Db
	var res dto.VoucherListRes
	query := entity.TbVoucher{
		ShopID: input.ShopID,
		Status: input.Status,
		Type:   input.Type,
	}
	db = db.Model(&entity.TbVoucher{}).Joins("left join tb_shop shop on tb_voucher.shop_id = shop.id")
	db = db.Select("tb_voucher.*, shop.name")
	db = db.Where(query)
	if input.Title != "" {
		db = db.Where("title like ?", "%"+input.Title+"%")
	}
	if input.ShopName != "" {
		db = db.Where("shop.name like ?", "%"+input.ShopName+"%")
	}
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *VoucherService) Update(input dto.UpdateVoucherReq) error {
	db := myGrom.Db
	voucher := entity.TbVoucher{
		Title:       input.Title,
		SubTitle:    input.SubTitle,
		Rules:       input.Rules,
		PayValue:    input.PayValue,
		ActualValue: input.ActualValue,
		Type:        input.Type,
		Status:      input.Status,
	}
	if err := db.Model(&entity.TbVoucher{}).Where("id", input.ID).Updates(&voucher).Error; err != nil {
		return err
	}
	return nil
}

func (s *VoucherService) AddVoucher(input dto.CreateVoucherReq) error {
	db := myGrom.Db
	voucher := entity.TbVoucher{
		ShopID:      input.ShopID,
		Title:       input.Title,
		SubTitle:    input.SubTitle,
		Rules:       input.Rules,
		PayValue:    input.PayValue,
		ActualValue: input.ActualValue,
		Type:        input.Type,
		Status:      input.Status,
	}
	if err := db.Create(&voucher).Error; err != nil {
		return err
	}
	return nil
}

func (s *VoucherService) GetVoucher(input dto.GetVoucherInfoReq) (dto.VoucherOutModel, error) {
	db := myGrom.Db
	var res dto.VoucherOutModel
	if err := db.Model(&entity.TbVoucher{}).Where(input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *VoucherService) Remove(input dto.RemoveVoucherReq) error {
	db := myGrom.Db
	if err := db.Where("id", input.ID).Delete(&entity.TbVoucher{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *VoucherService) BatchRemove(input dto.BatchRemoveVoucherReq) error {
	db := myGrom.Db
	ids := str.IntStrToArray(input.IDS, ",")
	if err := db.Where("id", ids).Delete(&entity.TbVoucher{}).Error; err != nil {
		return err
	}
	return nil
}
