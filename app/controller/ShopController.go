package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type ShopController struct {
	shopService service.ShopService
}

func (c *ShopController) ShopList(ctx *gin.Context) {
	res, err := c.shopService.ShopList()
	if err != nil {
		myGin.Failure(ctx, "商店列表查询失败")
		return
	}
	myGin.Success(ctx, res)
}
