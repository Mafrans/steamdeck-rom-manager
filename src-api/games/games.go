package games

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"mafrans/steamdeck-rom-manager/utils"
	"os"
	"path/filepath"
)

type GameMeta struct {
	Game *Game  `json:"game"`
	File string `json:"file"`
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
	gamedir := gamemeta.Game.Name
	if gamedir == "" {
		gamedir = fmt.Sprint(gamemeta.Game.Id)
	}

	metaFile := gamemeta.GetGamePath("metadata.json")
	json, _ := json.Marshal(gamemeta)

	os.WriteFile(metaFile, json, 0755)
}

func ById(id string) (GameMeta, bool) {
	metaFile := utils.GetConfigPath(filepath.Join("games", id), "metadata.json")
	content, _ := os.ReadFile(metaFile)

	var game GameMeta
	if json.Unmarshal(content, &game) == nil {
		return game, true
	}
	return game, false
}

func (gamemeta *GameMeta) GetGamePath(file string) string {
	return utils.GetConfigPath(filepath.Join("games", gamemeta.Game.Name), file)
}
