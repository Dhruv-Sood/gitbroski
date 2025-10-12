// Package version holds the gitbroski CLI version.
package version

// Version is the current version of the CLI.
// It can be overridden at build time using:
//
//	go build -ldflags "-X gitbroski/internal/version.Version=1.3.0" -o gitbroski ./cmd
var Version = "dev"
