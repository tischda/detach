package main

import (
	"flag"
	"fmt"
	"os"
)

// Set at build time via -ldflags "-X main.AppName=... -X main.AppVersion=... -X main.BuildDate=... -X main.GitCommit=..."
var AppName string
var AppVersion string
var BuildDate string
var GitCommit string

var flagHelp = flag.Bool("help", false, "displays this help message")
var flagVersion = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(flagHelp, "h", false, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+AppName+` [ version | --version | --help ]

Description here.

OPTIONS:

  -h, -help
        display this help message
  -v, -version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+AppName+`
    Example output here.`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || *flagVersion {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", AppName, AppVersion, BuildDate, GitCommit)
		return
	}

	if *flagHelp {
		flag.Usage()
		return
	}

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Do what you have to do here.
	// process()
}
