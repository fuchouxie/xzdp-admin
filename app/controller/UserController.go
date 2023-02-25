package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type UserController struct {
	userService service.UserService
}

func (c *UserController) UserList(ctx *gin.Context) {
	var req dto.UserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.userService.UserList(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}
