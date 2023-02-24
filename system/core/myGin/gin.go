package myGin

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"xzdp-admin/system/code"
	"xzdp-admin/system/utils/config"
)

type ResponseParams struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var myEngine *gin.Engine

// 初始化
func InitGin() *gin.Engine {

	file, _ := os.Create("./logs/web.log")
	gin.DefaultWriter = io.MultiWriter(file)
	//日志同时输出到控制台
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	myEngine = gin.Default()
	return myEngine
}

func Start() {
	err := myEngine.Run(":" + config.GetConfigString("port"))
	if err != nil {
		panic(err)
		return
	}
}

//发送成功请求

func Success(ctx *gin.Context, data interface{}) {
	DoResponse(ctx, ResponseParams{
		Code: code.CODE_SUCCESS,
		Data: data,
		Msg:  "",
	})
}

//发送失败请求

func Failure(ctx *gin.Context, msg string) {
	DoResponse(ctx, ResponseParams{
		Code: code.CODE_FAIL,
		Data: nil,
		Msg:  msg,
	})
}

func DoResponse(ctx *gin.Context, params ResponseParams) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": params.Code,
		"data": params.Data,
		"msg":  params.Msg,
	})
}

func SetupRouterOfController(e *gin.Engine, routerPath string, ctrl interface{}) {
	v := reflect.ValueOf(ctrl)
	t := reflect.TypeOf(ctrl)
	for i := 0; i < v.NumMethod(); i++ {
		fn := t.Method(i).Name
		isGet := ShouldOnlyGet(fn)
		if fc, ok := v.Method(i).Interface().(func(ctx *gin.Context)); ok {
			if isGet {
				e.GET(routerPath+fn, fc)
			} else {
				e.POST(routerPath+fn, fc)
			}
		}
	}
}

// 判断该接口是否应该为GET请求
func ShouldOnlyGet(fn string) bool {
	fnName := strings.ToLower(fn)
	res := strings.Contains(fnName, "get")
	if res {
		return true
	}
	res = strings.Contains(fnName, "list")
	if res {
		return true
	}
	return false
}
