package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type ShopController struct {
	shopService service.ShopService
}

func (c *ShopController) ShopList(ctx *gin.Context) {
	var req dto.ShopListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.shopService.ShopList(req)
	if err != nil {
		myGin.Failure(ctx, "商店列表查询失败")
		return
	}
	myGin.Success(ctx, res)
}

func (c *ShopController) Create(ctx *gin.Context) {
	var req dto.CreateShopReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	if err := c.shopService.Create(req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *ShopController) Update(ctx *gin.Context) {
	var req dto.UpdateShopReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	if err := c.shopService.Update(req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *ShopController) Remove(ctx *gin.Context) {
	var req dto.RemoveShopReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	if err := c.shopService.Remove(req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *ShopController) BatchRemove(ctx *gin.Context) {
	var req dto.BatchRemoveShopReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	if err := c.shopService.BatchRemove(req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
