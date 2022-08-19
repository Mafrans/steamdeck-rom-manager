package games

import (
	"encoding/json"
	"log"
	"mafrans/steamdeck-rom-manager/src/utils"
	"os"
	"path/filepath"
)

type Game struct {
	Id   string
	Name string
}

func All() []Game {
	gameDir := utils.GetConfigPath("games", "")
	gameFolders, err := os.ReadDir(gameDir)
	if err != nil {
		log.Fatal(err)
	}

	games := make([]Game, 0)
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

func (game *Game) Save() {
	metaFile := utils.GetConfigPath(filepath.Join("games", game.Id), "metadata.json")
	json, _ := json.Marshal(game)

	os.WriteFile(metaFile, json, 0755)
}

func ById(id string) (Game, bool) {
	metaFile := utils.GetConfigPath(filepath.Join("games", id), "metadata.json")
	content, _ := os.ReadFile(metaFile)

	var game Game
	if json.Unmarshal(content, &game) == nil {
		return game, true
	}
	return game, false
}
