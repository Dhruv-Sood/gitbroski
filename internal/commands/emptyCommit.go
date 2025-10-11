package commands

import (
	"gitbroski/internal/services"
	"gitbroski/utils/logger"
)

func init() {
	Register("empty", emptyCommit)
}

func emptyCommit(args ...string) {
	if len(args) == 1 {
		logger.Error("Please Enter commit message, gitbroski empty commit <your-msg>")
		return
	}
	services.EmptyCommit(args[1])
}
