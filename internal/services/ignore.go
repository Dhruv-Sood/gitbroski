// Package services provides core business logic services.
package services

import (
	"gitbroski/assets/ignore"
	"gitbroski/utils/git"
	"gitbroski/utils/logger"
	"os"
	"path/filepath"
)

func Ignore(lang string) {
	rootDir := git.GetRoot()
	if rootDir == "" {
		return
	}

	gitignorePath := filepath.Join(rootDir, ".gitignore")
	switch lang {
	case "python":
		pythonIgnoreTemplate := ignore.Python()
		err := os.WriteFile(gitignorePath, []byte(pythonIgnoreTemplate), 0o600)
		if err != nil {
			logger.Error("Error writing .gitignore: " + err.Error())
		}
	case "go":
		goIgnoreTemplate := ignore.Golang()
		err := os.WriteFile(gitignorePath, []byte(goIgnoreTemplate), 0o600)
		if err != nil {
			logger.Error("Error writing .gitignore: " + err.Error())
		}
	default:
		err := os.WriteFile(gitignorePath, []byte{}, 0o600)
		if err != nil {
			logger.Error("Error creating empty .gitignore: " + err.Error())
		}
	}
}
