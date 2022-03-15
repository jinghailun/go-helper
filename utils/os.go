package utils

import (
	"fmt"
	"os"
)

func SavePid() int {
	excuPath,_ := os.Getwd()
	pid := os.Getpid()
	pidFile := excuPath + "/server.pid"
	WriteFile(pidFile, fmt.Sprintf("%d", pid))

	return pid
}
