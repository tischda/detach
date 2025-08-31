//go:build windows

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

var modkernel32 = syscall.NewLazyDLL("kernel32.dll")
var procExpandEnvironmentStringsW = modkernel32.NewProc("ExpandEnvironmentStringsW")

// getUserAndSystemEnv retrieves the current environment and overrides possibly stale values with
// USER and SYSTEM environment variables from the Windows registry. The "Path" variable is merged
// and SYSTEM variables take precedence over USER variables (which is the standard in Windows).
//
// It returns a slice of strings in "key=value" format containing all environment variables found,
// or an error if the registry cannot be accessed.
func getUserAndSystemEnv() ([]string, error) {
	envMap := make(map[string]string)

	// Obtain COMPUTERNAME, SYSTEMDRIVE, USERPROFILE, etc. from current environment
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "=") {
			continue // Skip entries starting with "="
		}
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	// Read SYSTEM environment vars
	sysReg, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.READ)
	if err == nil {
		defer sysReg.Close()
		sysEnv, _ := sysReg.ReadValueNames(0)
		for _, name := range sysEnv {
			val, _, _ := sysReg.GetStringValue(name)
			envMap[name] = val
		}
	}

	// Read USER environment vars
	userReg, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.READ)
	if err == nil {
		defer userReg.Close()
		userEnv, _ := userReg.ReadValueNames(0)
		for _, name := range userEnv {
			val, _, _ := userReg.GetStringValue(name)
			if name == "Path" {
				// Append USER Path to SYSTEM Path (system first, then user)
				envMap[name] = envMap[name] + ";" + val
			} else {
				envMap[name] = val
			}
		}
	}

	// Construct env []string in "key=value" format, expanding variables as we go
	env := make([]string, 0, len(envMap))
	for k, v := range envMap {
		v = expandVariable(v)
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env, nil
}

// expandVariable returns the resolved value of environment variables containing
// other variables such as %APPDATA% or %USERPROFILE%.
//
// Parameters:
//   - v: The environment variable to expand.
func expandVariable(v string) string {
	src, err := syscall.UTF16PtrFromString(v)
	if err != nil {
		log.Fatalln("String with NULL passed to StringToUTF16Ptr")
	}
	buf := make([]uint16, 32767) // Maximum environment variable size on Windows
	dst := &buf[0]
	size := uintptr(len(buf))

	n, _, _ := procExpandEnvironmentStringsW.Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(dst)),
		size,
	)
	if n != 0 && n <= size {
		v = syscall.UTF16ToString(buf[:n-1])
	}
	return v
}
