package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gwx/api"
)

func InitRouter(router *gin.Engine) {

	router.GET("/vw1/:tag", api.CheckSignature)
}
