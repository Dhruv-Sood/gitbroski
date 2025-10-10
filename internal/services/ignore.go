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
			return
		}
		logger.Success("Python .gitignore successfully created at " + gitignorePath)
	case "node", "nodejs", "js", "javascript":
		nodeIgnoreTemplate := ignore.NodeJS()
		err := os.WriteFile(gitignorePath, []byte(nodeIgnoreTemplate), 0644)
		if err != nil {
			logger.Error("Error writing .gitignore: " + err.Error())
			return
		}
		logger.Success("Node.js .gitignore successfully created at " + gitignorePath)
	default:
		err := os.WriteFile(gitignorePath, []byte{}, 0o600)
		if err != nil {
			logger.Error("Error creating empty .gitignore: " + err.Error())
			return
		}
		if lang != "" {
			logger.Warning("Unknown language '" + lang + "'. Created empty .gitignore at " + gitignorePath)
		}
	}
}
