package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnvFromProjectRoot() error {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return err
	}
	return godotenv.Load(filepath.Join(projectRoot, ".env"))
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf(".env not found in any parent dir")
		}
		dir = parent
	}
}
