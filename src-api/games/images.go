package games

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadCoverArt(meta GameMeta) string {
	url := fmt.Sprintf(
		"https://github.com/libretro-thumbnails/%s/raw/master/Named_Boxarts/%s.png",
		strings.ReplaceAll(meta.Game.Console, " ", "_"),
		meta.Game.Name,
	)
	resp, err := http.Get(url)

	log.Println(url)

	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return ""
	}

	if body == nil {
		log.Println("[IMAGES] Body is null")
		return ""
	}

	// log.Println(string(body))

	path := meta.GetGamePath("cover.png")
	err = os.WriteFile(path, body, fs.ModeAppend)

	if err != nil {
		log.Printf("[IMAGES] Couldn't save cover image: %b", err)
		return ""
	}
	return path
}
