package commands

import "gitbroski/internal/services"

func init() {
	// Register canonical command and common flag aliases
	Register("version", version)
	Register("--version", version)
	Register("-v", version)
}

func version(_ ...string) {
	services.PrintVersion()
}
