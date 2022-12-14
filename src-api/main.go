package main

import (
	"fmt"
	"mafrans/steamdeck-rom-manager/games"
	"mafrans/steamdeck-rom-manager/upload"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Authorization", "Origin", "X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override", "Content-Type", "Upload-Length", "Upload-Offset", "Tus-Resumable", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"}
	corsConfig.MaxAge = 86400

	app.Use(cors.New(corsConfig))

	upload.CreateUploaderRoutes(app)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"text": "Hello World!",
		})
	})

	app.GET("/games", func(ctx *gin.Context) {
		games := games.All()
		fmt.Println(games)

		ctx.JSONP(http.StatusOK, games)
	})

	app.GET("/games/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		game, ok := games.ById(id)

		if ok {
			ctx.JSONP(http.StatusOK, game)
		} else {
			ctx.Status(http.StatusNotFound)
			ctx.Done()
		}
	})

	app.GET("/games/:id/cover", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		game, ok := games.ById(id)

		if ok {
			ctx.File(game.Artwork.Cover)
			ctx.Done()
		} else {
			ctx.Status(http.StatusNotFound)
			ctx.Done()
		}
	})

	app.DELETE("/games", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
		ctx.Done()
	})

	app.Run("0.0.0.0:3123")
}
