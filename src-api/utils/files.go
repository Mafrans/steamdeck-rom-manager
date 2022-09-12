package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
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

func HashFileSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha1.New()

	// Copy the file into the hash interface
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Get the first 20 bytes of the hash
	hashInBytes := hash.Sum(nil)[:20]
	return hex.EncodeToString(hashInBytes), nil

}
