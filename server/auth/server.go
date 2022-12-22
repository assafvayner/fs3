package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/assafvayner/fs3/protos/authservice"
	"github.com/assafvayner/fs3/server/auth/config"
	"github.com/assafvayner/fs3/server/auth/service"
	"github.com/assafvayner/fs3/server/shared/loggerutils"
	"github.com/assafvayner/fs3/server/shared/tlsutils"
)

func main() {
	logger := loggerutils.InitLogger("authservice")

	ln, err := net.Listen("tcp", fmt.Sprint(":", config.GetPort()))
	if err != nil {
		logger.Fatalln(err)
	}
	defer ln.Close()

	tlsCredentials, err := tlsutils.GetServerTLSCredentials()
	if err != nil {
		logger.Fatalln(err)
	}

	server := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	authservice.RegisterAuthServer(server, service.NewAuthServiceHandler(logger))

	err = server.Serve(ln)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start server serving")
		os.Exit(1)
	}
}
