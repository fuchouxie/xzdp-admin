package service

import (
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
	"xzdp-admin/system/lib/str"
)

type ShopService struct {
}

func (s *ShopService) ShopList(input dto.ShopListReq) (dto.ShopListRes, error) {
	db := myGrom.Db
	var res dto.ShopListRes
	query := entity.TbShop{
		ID:     input.ID,
		TypeID: input.TypeID,
		Score:  input.Score,
	}
	db = db.Model(entity.TbShop{}).Where(query)
	if input.Name != "" {
		db = db.Where("name like ?", "%"+input.Name+"%")
	}
	if input.Area != "" {
		db = db.Where("area like ?", "%"+input.Area+"%")
	}
	if input.Address != "" {
		db = db.Where("address like ?", "%"+input.Address+"%")
	}
	err := db.Scopes(myGrom.Paginate(input.PageLimit)).Scan(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *ShopService) Update(input dto.UpdateShopReq) error {
	db := myGrom.Db
	shop := entity.TbShop{
		ID:        input.ID,
		Name:      input.Name,
		TypeID:    input.TypeID,
		Images:    input.Images,
		Area:      input.Area,
		Address:   input.Address,
		X:         input.X,
		Y:         input.Y,
		AvgPrice:  input.AvgPrice,
		Sold:      input.Sold,
		Comments:  input.Comments,
		Score:     input.Score,
		OpenHours: input.OpenHours,
	}
	if err := db.Model(&entity.TbShop{}).Where("id", input.ID).Updates(&shop).Error; err != nil {
		return err
	}
	return nil
}

func (s *ShopService) Remove(input dto.RemoveShopReq) error {
	db := myGrom.Db
	if err := db.Model(&entity.TbShop{}).Delete("id", input.ID).Error; err != nil {
		return err
	}
	return nil
}

func (s *ShopService) BatchRemove(input dto.BatchRemoveShopReq) error {
	db := myGrom.Db
	ids := str.IntStrToArray(input.IDS, ",")
	if err := db.Model(&entity.TbShop{}).Delete("id", ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *ShopService) Create(input dto.CreateShopReq) error {
	db := myGrom.Db
	shop := entity.TbShop{
		Name:      input.Name,
		TypeID:    input.TypeID,
		Images:    input.Images,
		Area:      input.Area,
		Address:   input.Address,
		X:         input.X,
		Y:         input.Y,
		AvgPrice:  input.AvgPrice,
		Sold:      input.Sold,
		Comments:  input.Comments,
		Score:     input.Score,
		OpenHours: input.OpenHours,
	}
	if err := db.Model(&entity.TbShop{}).Create(&shop).Error; err != nil {
		return err
	}
	return nil
}
