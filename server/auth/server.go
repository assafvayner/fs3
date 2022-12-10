package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/service"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/loggerutils"
)

func main() {
	logger := loggerutils.InitLogger("authservice")

	ln, err := net.Listen("tcp", fmt.Sprint(":", config.GetPort()))
	if err != nil {
		logger.Fatalln(err)
		os.Exit(1)
	}
	defer ln.Close()

	server := grpc.NewServer()
	authservice.RegisterAuthServer(server, service.NewAuthServiceHandler(logger))

	err = server.Serve(ln)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start server serving")
		os.Exit(1)
	}
}
