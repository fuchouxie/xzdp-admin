package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xzdp-admin/app/consts"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/gen/dal/entity"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
	"xzdp-admin/system/core/mySession"
)

type Controller struct {
	UserInfo     *entity.AdminUser
	adminService service.AdminService
}

func (c *Controller) CallBefore(ctx *gin.Context) error {
	c.UserInfo = &entity.AdminUser{}
	value, err := mySession.GetSession(ctx, "userId")
	if err != nil {
		return err
	}
	id, _ := strconv.ParseInt(value.(string), 10, 32)
	userInfo, err := c.adminService.GetAdminUser(dto.GetAdminInfoReq{ID: id})
	if err != nil {
		return err
	}
	c.UserInfo = &userInfo
	return nil
}

func (c *Controller) GetGlobalEnum(ctx *gin.Context) {
	globalEnum := map[string]interface{}{
		"enum_admin_type": consts.EnumAdminType,
	}
	myGin.Success(ctx, globalEnum)
}
