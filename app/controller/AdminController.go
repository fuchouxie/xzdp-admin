package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type AdminController struct {
	AdminService *service.AdminService
}

func (c *AdminController) Login(ctx *gin.Context) {
	var req dto.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.AdminService.Login(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *AdminController) Logout(ctx *gin.Context) {
	myGin.Success(ctx, nil)
}

func (c *AdminController) Register(ctx *gin.Context) {
	var req dto.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.AdminService.Register(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *AdminController) AdminList(ctx *gin.Context) {
	var req dto.AdminListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.AdminService.AdminList(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *AdminController) Update(ctx *gin.Context) {
	var req dto.UpdateAdminReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	if req.ID == 0 {
		myGin.Failure(ctx, "id不能为空")
		return
	}
	err := c.AdminService.UpdateInfo(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *AdminController) ChangePassword(ctx *gin.Context) {
	var req dto.ChangePasswordReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.AdminService.ChangePassword(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *AdminController) Remove(ctx *gin.Context) {
	var req dto.RemoveReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.AdminService.Remove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *AdminController) BatchRemove(ctx *gin.Context) {
	var req dto.BatchRemoveReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.AdminService.BatchRemove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
