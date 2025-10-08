package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/square/exit"
	"github.com/square/exoskeleton"
)

const Version = "0.0.1"

func main() {
	path, _ := os.Getwd()
	paths := []string{path + "/devcommands"}

    // Discover all the directories called `devcommands` from
    // where the command is executed, and up to the root.
    for path != "/" {
    	path = filepath.Dir(path)
    	devcommandsPath := filepath.Join(path, "devcommands")
    	if _, err := os.Stat(devcommandsPath); err == nil {
    		paths = append(paths, devcommandsPath)
    	}
    }

    // + Add the homebrew etc folders
	paths = append(paths, "/opt/homebrew/etc/devcommands") // Apple silicon
	paths = append(paths, "/usr/local/etc/devcommands") // Intel path

    // Display all path if DEV_CLI_DEBUG env variable is set
    if os.Getenv("DEV_CLI_DEBUG") != "" {
    	fmt.Fprintln(os.Stderr, "[DEBUG] Discovered command paths:")
    	for _, path := range paths {
    		fmt.Fprintln(os.Stderr, "[DEBUG]  -", path)
    	}
    }

	cli, err := exoskeleton.New(paths)
	if err != nil {
		panic(err)
	}

	// Identify the subcommand being invoked from the arguments.
	cmd, args, err := cli.Identify(os.Args[1:])
	if err != nil {
		panic(err)
	}

	// Execute the subcommand.
	err = cmd.Exec(cli, args, os.Environ())

	// Exit the program with the exit code the subcommand returned.
	os.Exit(exit.FromError(err))
}
