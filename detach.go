//go:build windows

package main

import (
	"log"
	"os/exec"

	"golang.org/x/sys/windows"
)

// detach starts a new process specified by cmdPath and cmdArgs in a detached state on Windows.
// The new process will not be attached to the current console and will run independently.
//
// The process will also inherit a new set of user and system environment variables.
//
// Parameters:
//   - cmdPath: The path to the executable to run.
//   - cmdArgs: The arguments to pass to the executable.
//
// The function logs fatal errors if process creation or release fails.
func detach(cmdPath string, cmdArgs []string) {
	env, err := getUserAndSystemEnv()
	if err != nil {
		log.Fatalln("Failed to get environment:", err)
	}

	cmd := exec.Command(cmdPath, cmdArgs...)
	cmd.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.CREATE_NEW_PROCESS_GROUP | windows.DETACHED_PROCESS,
	}
	cmd.Env = env

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	err = cmd.Start()
	if err != nil {
		log.Fatalln("Failed to start process:", err)
	}

	err = cmd.Process.Release()
	if err != nil {
		log.Fatalln("Failed to release process:", err)
	}

	log.Printf("Started detached process '%s' with PID %d\n", cmdPath, cmd.Process.Pid)
}
