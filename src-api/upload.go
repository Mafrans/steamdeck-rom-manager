package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

//go:embed uploader/assets/*
var uploaderAssets embed.FS

//go:embed uploader/index.html
var uploaderIndex string

// CreateUploaderRoutes creates the gin routes for tusd
func CreateUploaderRoutes(relativePath string, app *gin.Engine) {
	handler := createHandler()

	go handleUploadEvents(handler)

	app.GET(relativePath, func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", []byte(uploaderIndex))
	})

	app.GET(relativePath + "/assets/:asset", func(ctx *gin.Context) {
		asset := ctx.Param("asset")
		ctx.FileFromFS(path.Join("uploader/assets", asset), http.FS(uploaderAssets))
	})

	// Handle file upload routes
	app.Any(relativePath + "/files/*file", gin.WrapH(http.StripPrefix("/files/", handler.Middleware(handler))))
}

func createHandler() *tusd.Handler {
	composer := tusd.NewStoreComposer()
	store := filestore.FileStore{
		Path: PrepareFile(GetTmpDir(), ""),
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
		meta := IdentifyGame(uploadPath)
		if meta.Game == nil {
			log.Println("Upload failed: Failed to identify ROM details")
			return
		}

		gameDir, err := meta.GetGameDir();
		if err != nil {
			log.Println("Upload failed: Could not access game directory")
			return
		}

		meta.File = filepath.Join(gameDir, event.Upload.MetaData["filename"])
		os.Rename(uploadPath, meta.File)

		meta.Artwork.Cover = filepath.Join(gameDir, "cover.png")
		if err = DownloadCoverArt(meta.Game, meta.Artwork.Cover); err != nil {
			log.Println("Couldn't download cover art:", err)
		}

		meta.Save()

		// Delete temporary files
		os.Remove(uploadPath)
		os.Remove(uploadPath + ".info")
	}
}
