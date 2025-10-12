package commands

import (
	"gitbroski/internal/version"
	"gitbroski/utils/logger"
)

func init() {
	Register("version", Version)
	Register("--version", Version)
	Register("-v", Version)
}

// Version prints the current gitbroski version.
func Version(_ ...string) {
	logger.Text("gitbroski " + version.Version)
}
