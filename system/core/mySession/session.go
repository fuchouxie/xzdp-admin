package mySession

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

var SessionMap map[string]*sessions.Session

func init() {
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
	if session == nil {
		return nil, errors.New("未登录")
	}
	value := session.Values[key]
	return value.(string), nil
}

func SaveSession(ctx *gin.Context, key string, value string) {
	//1.生成uuid
	id := uuid.NewV4()
	_uuid := id.String()
	//2.初始化该用户的session
	SessionMap[_uuid] = &sessions.Session{
		ID:     _uuid,
		Values: make(map[interface{}]interface{}),
	}
	//3.设置cookie
	SessionMap[_uuid].Values[key] = value
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "session-id",
		Value:   _uuid,
		Path:    "/",
		Domain:  "",
		Expires: time.Now().Add(24 * time.Hour),
	})
}
