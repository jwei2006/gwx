package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gwx/config"
	"github.com/jwei2006/gwx/system"
)

func tagFilter(c *gin.Context)  {
	tag := c.Param("tag")
	if tag == ""{
		return
	}
	has, _ := system.GetContext(tag)
	if !has{
		return
	}
	c.Set(config.Key_Tag, tag)
}