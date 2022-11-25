package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/operations"
)

func main() {
	// defining parser to parse args to run operation
	parser := argparse.NewParser("fs3", "fs3 command line tool to copy/remote/get files from remote server")

	// define commands
	cpCmd := parser.NewCommand("cp", "Copy local file to remote server")
	rmCmd := parser.NewCommand("rm", "Remove file from remote server")
	getCmd := parser.NewCommand("get", "Retreive file from remote server")

// TODO: change local paths to FilePositionals: https://pkg.go.dev/github.com/akamensky/argparse#Command.FilePositional

	// cpCmd args
	cpLocalFile := cpCmd.FilePositional(os.O_RDONLY, 0444, &argparse.Options{
		Help: "local file to copy to server",
	})
	cpRemotePath := cpCmd.StringPositional(&argparse.Options{
		Help:     "remote path to put file on server, relative to root",
	})

	// rmCmd args
	rmRemotePath := rmCmd.StringPositional(&argparse.Options{
		Help:     "remote path to file you want to delete from server",
	})

	// getCmd args
	getRemotePath := getCmd.StringPositional(&argparse.Options{
		Help:     "remote path to file to get from server, relative to root",
	})
	getLocalFile := getCmd.FilePositional(os.O_WRONLY | os.O_CREATE, 0666, &argparse.Options{
		Help:     "local file to write to",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	// operations in charge of closing open files
	if cpCmd.Happened() {
		operations.Copy(cpLocalFile, *cpRemotePath)
	} else if rmCmd.Happened() {
		operations.Remove(*rmRemotePath)
	} else if getCmd.Happened() {
		operations.Get(*getRemotePath, getLocalFile)
	} else {
		err = fmt.Errorf("bad arguments for fs3")
		fmt.Fprintf(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}
}

