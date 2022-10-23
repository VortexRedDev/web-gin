package main

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/router"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := router.Router()
	r.Run()
}
