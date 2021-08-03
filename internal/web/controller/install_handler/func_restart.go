package install_handler

import (
	"bytes"
	"os/exec"
	"runtime"

	"exams-api/internal/pkg/core"
)

func (h *handler) Restart() core.HandlerFunc {
	return func(c core.Context) {
		shellPath := "./scripts/restart.sh"

		// runtime.GOOS = linux or darwin
		command := exec.Command("/bin/bash", "-c", shellPath)

		if runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/C", shellPath)
		}

		var stderr bytes.Buffer
		command.Stderr = &stderr
		outPutString := ""

		output, err := command.Output()
		if err != nil {
			outPutString += stderr.String()
			c.Payload(outPutString)
			return
		}

		outPutString += string(output)

		c.Payload(outPutString)
	}
}
