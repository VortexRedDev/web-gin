package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHomeRouter(r *gin.RouterGroup) {
	home := r.Group("/home")
	home.GET("/1", func(c *gin.Context) {
		c.String(http.StatusOK, "home1")
	})
	home.GET("/2", func(c *gin.Context) {
		c.String(http.StatusOK, "home2")
	})
}
