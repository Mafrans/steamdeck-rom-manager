package games

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"mafrans/steamdeck-rom-manager/utils"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type GameMeta struct {
	Game    *Game  `json:"game"`
	File    string `json:"file"`
	Artwork struct {	
		Cover string `json:"cover"`
	} `json:"artwork"`
}

func All() []GameMeta {
	gameDir := utils.GetConfigPath("games", "")
	gameFolders, err := os.ReadDir(gameDir)
	if err != nil {
		log.Fatal(err)
	}

	games := make([]GameMeta, 0)
	for _, folder := range gameFolders {
		if folder.IsDir() {
			game, ok := ById(folder.Name())
			if ok {
				games = append(games, game)
			}
		}
	}

	return games
}

func (gamemeta *GameMeta) Save() {
	metaFile := gamemeta.GetGamePath("metadata.toml")

	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	enc.Encode(gamemeta)

	os.WriteFile(metaFile, buf.Bytes(), 0755)
}

func ById(id string) (GameMeta, bool) {
	metaFile := utils.GetConfigPath(filepath.Join("games", id), "metadata.toml")
	content, _ := os.ReadFile(metaFile)

	var game GameMeta
	if _, err := toml.Decode(string(content), &game); err == nil {
		return game, true
	}
	return game, false
}

func (gamemeta *GameMeta) GetGamePath(file string) string {
	return utils.GetConfigPath(
		filepath.Join("games", fmt.Sprintf("%d", *gamemeta.Game.Id)),
		file,
	)
}
