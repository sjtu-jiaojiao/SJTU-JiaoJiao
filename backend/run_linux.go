// +build !windows

package main

import (
	"os/exec"
	"syscall"
)

func kill(cmd *exec.Cmd) {
	_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
}
