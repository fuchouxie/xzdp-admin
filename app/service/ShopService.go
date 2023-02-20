package service

import (
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
)

type ShopService struct {
}

func (s *ShopService) ShopList() ([]entity.TbShop, error) {
	db := myGrom.Db
	var res []entity.TbShop
	err := db.Model(entity.TbShop{}).Where("1=1").Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
