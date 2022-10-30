package main

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/config"
	"github/xiaoda-ye/web-gin/log"
	"github/xiaoda-ye/web-gin/router"
	"net/http"
)

func main() {
	r := gin.New()
	//f, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//
	//r.Use(gin.LoggerWithWriter(io.MultiWriter(f, os.Stdout)))
	//r.Use(logerMiddleware())
	r.Use(gin.LoggerWithWriter(log.LoggerWriter()), gin.Recovery())
	r.StaticFS("/statics", http.Dir("statics"))
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")

	router.Router(r)
	r.Run(":" + config.Conf().Server.Port)

}
