package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"

	"github.com/assafvayner/fs3/cli-go/operations"
)

func main() {
	fmt.Fprintln(os.Stderr, "defunct until converted to use http frontend, grpc interface blocked to internal service communication")
	os.Exit(1)
	// defining parser to parse args to run operation
	parser := argparse.NewParser("fs3", "fs3 command line tool to copy/remote/get files from remote server")

	// define commands
	cpCmd := parser.NewCommand("cp", "Copy local file to remote server")
	rmCmd := parser.NewCommand("rm", "Remove file from remote server")
	getCmd := parser.NewCommand("get", "Retreive file from remote server")
	describeCmd := parser.NewCommand("describe", "Describe a file or directory")
	loginCmd := parser.NewCommand("login", "Login to use fs3 service in your user context")
	newuserCmd := parser.NewCommand("newuser", "Create a user to log in to fs3")

	// cpCmd args
	cpLocalFile := cpCmd.FilePositional(os.O_RDONLY, 0444, &argparse.Options{
		Help: "local file to copy to server",
	})
	cpRemotePath := cpCmd.StringPositional(&argparse.Options{
		Help: "remote path to put file on server, relative to root",
	})

	// rmCmd args
	rmRemotePath := rmCmd.StringPositional(&argparse.Options{
		Help: "remote path to file you want to delete from server",
	})

	// getCmd args
	getRemotePath := getCmd.StringPositional(&argparse.Options{
		Help: "remote path to file to get from server, relative to root",
	})
	getLocalFile := getCmd.FilePositional(os.O_WRONLY|os.O_CREATE, 0666, &argparse.Options{
		Help: "local file to write to",
	})

	// describeCmd args
	describePath := describeCmd.StringPositional(&argparse.Options{
		Help: "path of file or directory to describe",
	})

	// loginCmd args
	loginUsername := loginCmd.String("u", "username", &argparse.Options{
		Help:     "your fs3 username",
		Required: false,
		Default:  "",
	})
	loginUseToken := loginCmd.Flag("t", "use-token", &argparse.Options{
		Help:     "should we use token to log in, if token invalid will still ask for password",
		Required: false,
		Default:  false,
	})

	// newuserCmd args
	newuserUsername := newuserCmd.String("u", "username", &argparse.Options{
		Help:     "your fs3 username",
		Required: true,
		Default:  "",
	})
	newuserPassword := newuserCmd.String("p", "password", &argparse.Options{
		Help:     "your desired fs3 password",
		Required: true,
		Default:  "",
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
		if argparse.IsNilFile(getLocalFile) {
			fmt.Fprintln(os.Stderr, "invalid destination file")
			os.Exit(1)
		}
		operations.Get(*getRemotePath, getLocalFile)
	} else if describeCmd.Happened() {
		operations.Describe(*describePath)
	} else if loginCmd.Happened() {
		operations.Login(*loginUsername, *loginUseToken)
	} else if newuserCmd.Happened() {
		operations.NewUser(*newuserUsername, *newuserPassword)
	} else {
		err = fmt.Errorf("bad arguments for fs3")
		fmt.Fprintln(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}
}
