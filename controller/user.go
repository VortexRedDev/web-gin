package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("/1", func(c *gin.Context) {
		c.String(http.StatusOK, "user1")
	})
	user.GET("/2", func(c *gin.Context) {
		//c.String(http.StatusOK, "user2")
		c.JSON(200, gin.H{"message": "poing"})
	})
}
