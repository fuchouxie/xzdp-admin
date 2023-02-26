package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/consts"
	"xzdp-admin/system/core/myGin"
)

type Controller struct {
}

func (c *Controller) GetGlobalEnum(ctx *gin.Context) {
	globalEnum := map[string]interface{}{
		"enum_admin_type": consts.EnumAdminType,
	}
	myGin.Success(ctx, globalEnum)
}
