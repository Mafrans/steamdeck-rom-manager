package upload

import (
	"embed"
	"log"
	"mafrans/steamdeck-rom-manager/games"
	"mafrans/steamdeck-rom-manager/utils"
	"net/http"
	"os"
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

		uploadPath := event.Upload.Storage["Path"]
		meta := games.Identify(uploadPath)
		meta.File = meta.GetGamePath(event.Upload.MetaData["filename"])
		
		if meta.Game != nil {
			os.Rename(uploadPath, meta.File)

			cover := meta.GetGamePath("cover.png")
			err := games.DownloadCoverArt(meta.Game, cover);
			if err != nil {
				println(err)
			} else {
				meta.Artwork.Cover = cover
			}

			meta.Save()

			// Delete temporary files
			os.Remove(uploadPath)
			os.Remove(uploadPath + ".info")
		}
	}
}
