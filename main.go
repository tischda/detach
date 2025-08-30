package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: detach <command_path> [args...]")
		os.Exit(1)
	}

	cmdPath := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := exec.Command(cmdPath, cmdArgs...)
	cmd.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.CREATE_NEW_PROCESS_GROUP | windows.DETACHED_PROCESS,
	}

	cmd.Env = os.Environ()

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process:", err)
		return
	}

	err = cmd.Process.Release()
	if err != nil {
		fmt.Println("Failed to release process:", err)
		return
	}

	fmt.Printf("Started detached process '%s' with PID %d\n", cmdPath, cmd.Process.Pid)
}
