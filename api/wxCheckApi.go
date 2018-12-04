package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jwei2006/gwx/context"
	"github.com/jwei2006/gwx/util"
)

func CheckSignature(c *gin.Context)  {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	s := util.Signature(timestamp, nonce, wechatCtx.Token)
	fmt.Println(signature == s)
	c.String(http.StatusOK, echostr)
}