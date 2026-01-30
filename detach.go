//go:build windows

package main

import (
	"errors"
	"fmt"
	"os/exec"

	"golang.org/x/sys/windows"
)

// detach starts a new process specified by cmd in a detached state on Windows.
// The new process will not be attached to the current console and will run independently.
//
// The process will also inherit a new set of user and system environment variables.
//
// Parameters:
//   - cmd: The executable to run and its arguments as a slice of strings.
//
// The function returns the process ID of the new process or an error if the process creation fails.
func detach(cmd []string) (int, error) {
	if len(cmd) == 0 {
		return 0, errors.New("command array is empty")
	}
	c := exec.Command(cmd[0], cmd[1:]...)

	c.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.DETACHED_PROCESS,
	}

	// prepare environment for process
	env, err := getUserAndSystemEnv()
	if err != nil {
		return 0, fmt.Errorf("failed to get environment: %w", err)
	}
	c.Env = env

	// start process
	err = c.Start()
	if err != nil {
		return 0, fmt.Errorf("failed to start command %v : %w", cmd, err)
	}
	return c.Process.Pid, nil
}
