package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type VoucherController struct {
	voucherService service.VoucherService
}

func (c *VoucherController) VoucherList(ctx *gin.Context) {
	var req dto.VoucherListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.voucherService.VoucherList(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *VoucherController) Create(ctx *gin.Context) {
	var req dto.CreateVoucherReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.voucherService.AddVoucher(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
func (c *VoucherController) Update(ctx *gin.Context) {
	var req dto.UpdateVoucherReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.voucherService.Update(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
func (c *VoucherController) GetVoucher(ctx *gin.Context) {
	var req dto.GetVoucherInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.voucherService.GetVoucher(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *VoucherController) Remove(ctx *gin.Context) {
	var req dto.RemoveVoucherReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.voucherService.Remove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *VoucherController) BatchRemove(ctx *gin.Context) {
	var req dto.BatchRemoveVoucherReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.voucherService.BatchRemove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
