package upload

import (
	"embed"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var fs embed.FS

//go:embed index.html
var index string

func CreateUploaderRoutes(app *gin.Engine) {
	app.GET("/upload", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte(index))
	})

	app.GET("/upload/assets/:asset", func(ctx *gin.Context) {
		asset := ctx.Param("asset")
		ctx.FileFromFS(path.Join("assets", asset), http.FS(fs))
	})
}
