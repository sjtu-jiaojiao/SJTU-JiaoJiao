// +build windows

package main

import (
	"os/exec"
	"strconv"
)

func kill(cmd *exec.Cmd) {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
	_ = kill.Start()
}
