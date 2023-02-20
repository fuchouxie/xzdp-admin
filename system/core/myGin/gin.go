package myGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
