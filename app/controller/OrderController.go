package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type OrderController struct {
	orderService service.OrderService
}

func (c *OrderController) OrderList(ctx *gin.Context) {
	var req dto.OrderListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.orderService.OrderList(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *OrderController) GetOrderInfo(ctx *gin.Context) {
	var req dto.GetOrderInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.orderService.GetOrder(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}
