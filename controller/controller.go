package controller

import (
	"github.com/gin-gonic/gin"
)

// ...
func AddRouter(r *gin.Engine) {
	router := r.Group("/api")
	addHomeRouter(router)
	addUserRouter(router)
}
