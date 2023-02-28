package mySession

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

var SessionMap map[string]*sessions.Session

func InitSession() {
	SessionMap = make(map[string]*sessions.Session)
}

func GetSession(ctx *gin.Context, key string) (interface{}, error) {
	var sessionId string
	cookie, err := ctx.Request.Cookie("session-id")
	if err != nil {
		return nil, err
	}
	sessionId = cookie.Value
	session := SessionMap[sessionId]
	value := session.Values[key]
	return value.(string), nil
}

func SaveSession(ctx *gin.Context, key string, value string) {
	//1.生成uuid
	uuid := "abcabc"
	//2.初始化该用户的session
	SessionMap[uuid] = &sessions.Session{
		ID:     uuid,
		Values: make(map[interface{}]interface{}),
	}
	//3.设置cookie
	SessionMap[uuid].Values[key] = value
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "session-id",
		Value:   uuid,
		Path:    "/",
		Domain:  "",
		Expires: time.Now().Add(24 * time.Hour),
	})
}
