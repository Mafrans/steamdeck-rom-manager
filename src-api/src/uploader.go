package main

import (
	"github.com/gin-gonic/gin"
)

func CreateUploaderRoutes(app *gin.Engine) {
	app.StaticFile("/upload", "./assets/uploader/index.html")
}
