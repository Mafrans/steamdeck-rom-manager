package upload

import (
	"embed"
	"log"
	"mafrans/steamdeck-rom-manager/utils"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

//go:embed assets/*
var fs embed.FS

//go:embed index.html
var index string

func CreateUploaderRoutes(app *gin.Engine) {
	handler := createHandler()

	go handleUploadEvents(handler)

	app.GET("/upload", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte(index))
	})

	app.GET("/upload/assets/:asset", func(ctx *gin.Context) {
		asset := ctx.Param("asset")
		ctx.FileFromFS(path.Join("assets", asset), http.FS(fs))
	})

	// Handle file upload routes
	app.Any("/files/*file", gin.WrapH(http.StripPrefix("/files/", handler.Middleware(handler))))
}

func createHandler() *tusd.Handler {
	composer := tusd.NewStoreComposer()
	store := filestore.FileStore{
		Path: utils.GetConfigPath("tmp", ""),
	}
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})

	if err != nil {
		log.Fatalf("[UPLOAD]: Unable to create handler: %s\n", err)
	}

	return handler
}

func handleUploadEvents(handler *tusd.Handler) {
	for {
		event := <-handler.CompleteUploads
		log.Printf("[UPLOAD]: Upload %s finished\n", event.Upload.ID)
		log.Println(event.Upload.Storage)

		hash, _ := utils.HashFileSHA1(event.Upload.Storage["Path"])
		log.Println(hash)
	}
}
