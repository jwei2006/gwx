package system

import "github.com/gin-gonic/gin"

func tagFilter(c *gin.Context)  {
	tag := c.Param("tag")
	if tag == ""{
		return
	}
	c.Set()
}