package version

// Version represents the current version of the gitbroski CLI.
// It is intended to be overridden at build time using:
//
//	-ldflags "-X 'gitbroski/internal/version.Version=1.3.0'"
var Version = "dev"
