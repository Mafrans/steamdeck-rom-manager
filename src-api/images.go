package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

// DownloadCoverArt downloads the cover art for a game from libretro-thumbnails and saves it at a path
func DownloadCoverArt(game *Game, path string) error {
	uri := fmt.Sprintf(
		"https://github.com/libretro-thumbnails/%s/raw/master/Named_Boxarts/%s.png",
		strings.ReplaceAll(game.Console, " ", "_"),
		game.Name,
	)

	resp, err := http.Get(uri)

	log.Println(uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if body == nil {
		return errors.New("Body is empty")
	}

	// Link to a different image
	if strings.HasSuffix(string(body), ".png") {
		return DownloadCoverArt(&Game{
			Name: strings.TrimSuffix(string(body), ".png"),
			Console: game.Console,
			Franchise: game.Franchise,
			CrcHash: game.CrcHash,
		}, path)
	}

	PrepareFile(path);
	return os.WriteFile(path, body, fs.ModeAppend)
}