package main

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.cs.washington.edu/assafv/fs3/server/frontend/server"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/loggerutils"
)

const PORT_ENV_VAR = "PORT"

func main() {
	logger := loggerutils.InitLogger("frontend")

	srv := server.NewFrontendServer(getPort(), logger)

	err := srv.Run()
	logger.Fatalln(err)
}

func getPort() int {
	portStr, gotten := os.LookupEnv(PORT_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "could not get frontend port")
		os.Exit(1)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not parse port number")
		os.Exit(1)
	}
	return port
}
