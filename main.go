package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] cmdPath cmdArgs...

Starts a new process specified by cmdPath and cmdArgs in a detached state on Windows.
The new process will not be attached to the current console and will run independently.

The process will inherit the current environment and update USER and SYSTEM environment
variables from the Windows registry.

OPTIONS:

  -?, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` alacritty.exe
    2025/08/30 22:32:16 Started detached process 'alacritty.exe' with PID 34568`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// process
	pid, err := detach(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Started detached process %v with PID %d\n", os.Args[1:], pid)
}
