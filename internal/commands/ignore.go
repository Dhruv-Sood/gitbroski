// Package commands provides CLI command handlers.
package commands

import (
	"gitbroski/internal/services"
	"strings"
)

func init() {
	Register("ignore", Ignore)
}

func Ignore(args ...string) {
	lang := ""
	if len(args) > 0 {
		lang = strings.ToLower(strings.TrimSpace(args[0]))
	}
	services.Ignore(lang)
}
