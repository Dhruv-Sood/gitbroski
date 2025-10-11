package services

import (
	"gitbroski/utils/logger"
	"os/exec"
)

func EmptyCommit(message string) {
	result := exec.Command("git", "commit","--allow-empty" ,"-m", message)
	output, err := result.Output()

	if err != nil {
		errmsg := "Something went wrong -> " + err.Error()
		logger.Error(errmsg)
	}
	logger.Text("Empty commit success")
	logger.Text(string(output))
}