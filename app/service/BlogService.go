package service

import (
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/system/core/myGrom"
	"xzdp-admin/system/lib/str"
)

type BlogService struct {
}

func (s *BlogService) BlogList(input dto.BlogListReq) (dto.BlogListRes, error) {
	db := myGrom.Db
	var res dto.BlogListRes
	query := entity.TbBlog{
		ShopID: input.ShopID,
		UserID: input.UserID,
	}
	db = db.Model(&entity.TbBlog{})
	db = db.Joins("left join tb_shop shop on tb_blog.shop_id = shop.id")
	db = db.Joins("left join tb_user user on tb_blog.user_id = user.id")
	db = db.Select("tb_blog.*, shop.name, user.nick_name")
	db = db.Where(query)
	if input.Title != "" {
		db = db.Where("title like ?", "%"+input.Title+"%")
	}
	if err := db.Scopes(myGrom.Paginate(input.PageLimit)).Find(&res.List).Limit(-1).Offset(-1).Count(&res.Total).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *BlogService) GetBlog(input dto.GetBlogInfoReq) (dto.BlogOutModel, error) {
	db := myGrom.Db
	var res dto.BlogOutModel
	db = db.Model(&entity.TbBlog{})
	db = db.Joins("left join tb_shop shop on tb_blog.shop_id = shop.id")
	db = db.Joins("left join tb_user user on tb_blog.user_id = user.id")
	db = db.Select("tb_blog.*, shop.name, user.nick_name")
	if err := db.Where("tb_blog.id", input.ID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (s *BlogService) Remove(input dto.RemoveBlogReq) error {
	db := myGrom.Db
	if err := db.Where("id", input.ID).Delete(&entity.TbBlog{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *BlogService) BatchRemove(input dto.BatchRemoveBlogReq) error {
	db := myGrom.Db
	ids := str.IntStrToArray(input.IDS, ",")
	if err := db.Where("id", ids).Delete(&entity.TbBlog{}).Error; err != nil {
		return err
	}
	return nil
}
