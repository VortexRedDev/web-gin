package service

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/log"
)

func Auth(c *gin.Context) {
	log.Logger.Info("权限验证")
	//log.Panicln("权限验证")

	c.Next()
	//c.Abort()
}
