package main

import (
	"flag"
	"fmt"
	"os"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

var flagHelp = flag.Bool("help", false, "displays this help message")
var flagVersion = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(flagHelp, "h", false, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` cmdPath cmdArgs... | [ version | --version | --help ]

Starts a new process specified by cmdPath and cmdArgs in a detached state on Windows.
The new process will not be attached to the current console and will run independently.

The process will inherit the current environment and update USER and SYSTEM environment
variables from the Windows registry.

OPTIONS:

  -h, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` alacritty.exe
    2025/08/30 22:32:16 Started detached process 'alacritty.exe' with PID 34568`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || *flagVersion {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if *flagHelp {
		flag.Usage()
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// process
	detach(os.Args[1], os.Args[2:])
}
