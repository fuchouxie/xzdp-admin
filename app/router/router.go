package router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"xzdp-admin/app/controller"
)

func InitRouter(e *gin.Engine) {
	//setupShopRouter(e)
	setupRouter(e)
}
func setupRouterOfFunc(e *gin.Engine, routerPath string, handler gin.HandlerFunc) {
	e.GET(routerPath, handler)
	e.POST(routerPath, handler)

}
func setupRouterOfController(e *gin.Engine, routerPath string, ctrl interface{}) {
	v := reflect.ValueOf(ctrl)
	t := reflect.TypeOf(ctrl)
	for i := 0; i < v.NumMethod(); i++ {
		fn := t.Method(i).Name
		if fc, ok := v.Method(0).Interface().(func(ctx *gin.Context)); ok {
			e.GET(routerPath+fn, fc)
			e.POST(routerPath+fn, fc)
		}
	}
}

// store
/**
问题：这个注册路由的方法存在需要频繁改代码的问题，每当我们写了一个接口，就需要修改这个方法，这样可维护性不太好。
解决思路：将注册路由的粒度变粗，改为针对某个控制器的路由进行配置，这样仅仅需要在新增业务时才需要改代码。
*/
func setupShopRouter(e *gin.Engine) {
	storeController := new(controller.ShopController)
	rootPath := "/shop/"
	setupRouterOfFunc(e, rootPath+"StoreList", storeController.ShopList)

}

// 注册路由函数
// 以模块为单位将根路径以及所需注册的控制
func setupRouter(e *gin.Engine) {
	storeController := new(controller.ShopController)
	setupRouterOfController(e, "/shop/", storeController)
}
