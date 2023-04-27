package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// GameMeta is a set of metadata fields for a game
type GameMeta struct {
	Game    *Game  `json:"game"`
	File    string `json:"file"`
	Artwork struct {	
		Cover string `json:"cover"`
	} `json:"artwork"`
}

// AllGames scans the library and retrieves all games
func AllGames() []GameMeta {
	gamesDir, err := GetLibraryDir();
	if (err != nil) {
		log.Println("Couldn't find games directory")
		return []GameMeta{}
	}

	gameFolders, err := os.ReadDir(gamesDir)
	if err != nil {
		log.Println("Unable to read games directory:", err)
		return []GameMeta{}
	}

	games := make([]GameMeta, 0)
	for _, folder := range gameFolders {
		if folder.IsDir() {
			game, ok := GetGameByID(folder.Name())
			if ok {
				games = append(games, game)
			}
		}
	}

	return games
}

// Save saves a game's metadata into its metadata.toml file
func (gamemeta *GameMeta) Save() {
	gameDir, err := gamemeta.GetGameDir()
	if err != nil {
		log.Println("Couldn't save game: could not find game path")
	}

	metaFile := PrepareFile(gameDir, "metadata.toml")

	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	enc.Encode(gamemeta)

	os.WriteFile(metaFile, buf.Bytes(), 0755)
}

// GetGameByID retrieves a game by its id
func GetGameByID(id string) (GameMeta, bool) {
	var game GameMeta

	gamesDir, err := GetLibraryDir()
	if err != nil {
		log.Println("Couldn't find games directory")
		return game, false
	}

	metaFile := PrepareFile(gamesDir, id, "metadata.toml")
	content, _ := os.ReadFile(metaFile)

	if _, err := toml.Decode(string(content), &game); err == nil {
		return game, true
	}
	return game, false
}

// GetGameDir returns the directory a game is located in
func (gamemeta *GameMeta) GetGameDir() (string, error) {
	gamesDir, err := GetLibraryDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(gamesDir, fmt.Sprintf("%d", *gamemeta.Game.Id)), nil
}