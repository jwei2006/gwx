package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gwx/oauth"
	"github.com/jwei2006/gwx/system"
	"gwx/config"
	"gwx/util"
	"log"
	"net/http"
)

func Oauth(c *gin.Context)  {
	callback := c.Query("callback")
	if callback == ""{
		return
	}
	session := sessions.Default(c)
	session.Set(config.Key_Callback, callback)
	session.Save()
	scope := c.DefaultQuery("scope", "snsapi_base")
	state := c.DefaultQuery("state", "STATE")
	tag, _ := c.Get(config.Key_Tag)
	_, context := system.GetContext(tag.(string))
	newOauth := oauth.NewOauth(context)
	e := newOauth.Redirect(c.Writer, c.Request, "/vw1/" + tag.(string) + "/auth", scope, state)
	if e != nil{
		log.Println(e)
	}
}

func Auth(c *gin.Context)  {
	code := c.Query("code")
	if code == ""{
		return
	}
	_tag, _ := c.Get(config.Key_Tag)
	tag := _tag.(string)
	_, context := system.GetContext(tag)
	newOauth := oauth.NewOauth(context)
	accessToken, e := newOauth.GetUserAccessToken(code)
	if e != nil{
		log.Println(e)
		return
	}
	session := sessions.Default(c)
	bytes, e := util.AesEncrypt([]byte(accessToken.OpenID), []byte(config.AesKey))
	if e != nil{
		log.Println(e)
		return
	}
	session.Set(config.Key_Open, string(bytes))
	callback := session.Get(config.Key_Callback).(string)
	http.Redirect(c.Writer, c.Request, callback, http.StatusOK)
}