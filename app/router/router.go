package router

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/controller"
	"xzdp-admin/system/core/myGin"
)

func InitRouter(e *gin.Engine) {
	setupRouter(e)
}

// 注册路由函数
// 以模块为单位将根路径以及所需注册的控制
func setupRouter(e *gin.Engine) {
	storeController := new(controller.ShopController)
	myGin.SetupRouterOfController(e, "/shop/", storeController)
	adminController := new(controller.AdminController)
	myGin.SetupRouterOfController(e, "/admin/", adminController)
}
