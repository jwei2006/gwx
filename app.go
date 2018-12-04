package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gwx/system"
	"io"
	"gwx/config"
	"log"
	"os"
	"strconv"
)

func main() {
	gin.DisableConsoleColor()
	if !config.Debug {
		file, _ := os.Create(config.LogPath)
		gin.DefaultWriter = io.MultiWriter(file)
	}
	e := system.InitContext()
	if e != nil {
		log.Println(e)
		return
	}
	defer System.FinallyClose()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(config.CookieStore))
	router.Use(sessions.Sessions(config.SessionName, store))
	//router.Use(system.Around())
	system.InitRouter(router)

	router.Run(":" + strconv.Itoa(config.Port))
}
