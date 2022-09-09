package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func CheckFileExists(path string) bool {
	_, error := os.Stat(path)
	return !errors.Is(error, os.ErrNotExist)
}

func GetConfigPath(dir string, file string) string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		log.Println(err)
	}

	dirPath := filepath.Join(cfgDir, "steamdeck-rom-manager", dir)
	os.MkdirAll(dirPath, 0755)

	if file != "" {
		filePath := filepath.Join(dirPath, file)

		if !CheckFileExists(filepath.Join(dirPath, file)) {
			os.Create(filePath)
		}
		return filePath
	}

	return dirPath
}
