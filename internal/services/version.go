package services

import (
	"gitbroski/internal/version"
	"gitbroski/utils/logger"
)

// PrintVersion outputs the CLI name and current version.
func PrintVersion() {
	logger.Text("gitbroski " + version.Version)
}
