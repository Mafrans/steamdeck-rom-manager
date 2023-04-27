package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// AppDir is the app's directory name
const AppDir = "steamdeck-rom-manager"

// CheckFileExists checks if a file exists in the filesystem
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// GetConfigDir returns the config directory for the current OS
// on Windows, it is `C:/Users/<user>/AppData/roaming/<app>/config`
// on MacOS, it is `$HOME/Library/Application Support/<app>/config`
// on Linux, it is `$HOME/.config/<app>/config`
func GetConfigDir() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(config, AppDir, "config"), nil
}

// GetDataDir returns the data directory for the current OS
// on Windows, it is `C:/Users/<user>/<app>/data`
// on MacOS, it is `$HOME/Library/Application Support/<app>/data`
// on Linux, it is `$HOME/.local/share/<app>/data`
func GetDataDir() (string, error) {
	home, err := os.UserHomeDir();
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".local", "share", AppDir, "data")
	switch runtime.GOOS {
	case "windows":
		dir = filepath.Join(home, AppDir, "data")
	case "darwin":
		dir = filepath.Join(home, "Library", "Application Support", AppDir, "data")
	}

	return dir, nil
}

// GetLibraryDir returns the library directory
func GetLibraryDir() (string, error) {
	data, err := GetDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(data, "games"), nil
}

// GetTmpDir returns a suitable directory for storing temporary files
func GetTmpDir() (string) {
	return filepath.Join(os.TempDir(), AppDir)
}

// PrepareFile prepares a file for reading/writing
func PrepareFile(paths ...string) string {
	sep := string(filepath.Separator)

	// Join paths to make sure the complete path is split in its entirety
	parts := strings.Split(strings.Join(paths, sep), sep)
	file := parts[len(parts)-1]
	dir := strings.Join(parts[0:len(parts)-1], sep)

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Println(err)
	}

	if file != "" {
		fp := filepath.Join(dir, file)
		if !CheckFileExists(filepath.Join(dir, file)) {
			os.Create(fp)
		}
		return fp
	}

	return dir
}

// HashFileSHA1 opens a file and returns it's SHA1 hash
func HashFileSHA1(path string) (string, error) {
	file, err := os.Open(path)
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
