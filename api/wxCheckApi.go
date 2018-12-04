package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gwx/config"
	"github.com/jwei2006/gwx/system"
	"github.com/jwei2006/gwx/util"
	"net/http"
)

func CheckSignature(c *gin.Context)  {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	tag, _ := c.Get(config.Key_Tag)
	has, wechatCtx := system.GetContext(tag.(string))
	if !has{
		c.String(http.StatusOK, "")
		return
	}
	s := util.Signature(timestamp, nonce, wechatCtx.Token)
	fmt.Println(signature == s)
	c.String(http.StatusOK, echostr)
}