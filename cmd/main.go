package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hikmet-kibar/snip/cmd/app"
	"github.com/hikmet-kibar/snip/cmd/config"
)

// options are command-line options that are provided by the user.
type options struct {
	Directory string `short:"d" long:"directory" description:"Location of snips"`
}

func main() {
	// Add flags
	home, _ := os.UserHomeDir()
	defaultDir := filepath.Join(home, ".snips")
	// dirPtr := flag.String("directory", defaultDir, "Location of snips")

	// flag.Parse()

	lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	lsDirPtr := lsCmd.String(
		"directory",
		defaultDir,
		"Location of snips")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCmdDir := getCmd.String(
		"directory",
		defaultDir,
		"Location of snips")

	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	newCmdDir := newCmd.String(
		"directory",
		defaultDir,
		"Location of snips")

	// Convert to internal config

	//TODO: too few args provided
	// if len(os.Args) == 1 { }

	//TODO: help
	//TODO: default
	switch os.Args[1] {
	case "ls":
		lsCmd.Parse(os.Args[2:])
	case "new":
		newCmd.Parse(os.Args[2:])
	case "get":
		getCmd.Parse(os.Args[2:])
	default:
		fmt.Println("helping")
	}

	cfg := config.New()

	// TODO: Replace tilde with HOME in input
	if lsCmd.Parsed() {
		cfg.Directory = *lsDirPtr
		err := app.List(cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// TODO: Replace tilde with HOME in input
	if getCmd.Parsed() {
		cfg.Directory = *getCmdDir
		cfg.Snip = getCmd.Arg(0)
		err := app.Get(cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// TODO:
	//		- Replace tilde with HOME in input
	if newCmd.Parsed() {
		cfg.Directory = *newCmdDir
		cfg.Snip = newCmd.Arg(0)
		err := app.New(cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
