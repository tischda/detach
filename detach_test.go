//go:build windows

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestDetachIntegration(t *testing.T) {
	// Command to run in a detached process
	cmd := []string{"cmd", "/c", "C:\\Windows\\SysWOW64\\timeout.exe", "/T", "3", "/NOBREAK"}

	// Detach the process
	pid, err := detach(cmd)
	if err != nil {
		t.Fatalf("Failed to detach process: %v", err)
	}
	t.Logf("Started detached process with PID %d", pid)

	// Give the process a moment to start
	time.Sleep(1 * time.Second)

	// Check if the process is active by running tasklist
	out, err := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid)).Output()
	if err != nil {
		t.Fatalf("Failed to run tasklist: %v", err)
	}

	if !strings.Contains(string(out), "cmd") {
		t.Errorf("Process with PID %d is not active or not the correct process. Output: %s", pid, string(out))
	}
}
