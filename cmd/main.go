// Package main is the entry point for the gitbroski CLI application.
package main

import (
	"gitbroski/internal/commands"
	"gitbroski/utils/logger"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		logger.Banner("GitBroski")
		logger.Text("Welcome to GitBroski!\nThe best git helper tool.")
		return
	}

	cmd := os.Args[1]
	handler, exists := commands.Registry[cmd]

	if !exists {
		logger.Text("Command not found: " + cmd)
		return
	}
	handler(os.Args[2:]...)

}
