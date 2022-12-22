package main

import (
	"fmt"
	"net"
	"os"

	fs3 "github.com/assafvayner/fs3/protos/fs3"
	primarybackup "github.com/assafvayner/fs3/protos/primarybackup"
	"github.com/assafvayner/fs3/server/app/backup"
	"github.com/assafvayner/fs3/server/app/config"
	"github.com/assafvayner/fs3/server/app/primary"
	"github.com/assafvayner/fs3/server/app/primary/interceptor"
	"github.com/assafvayner/fs3/server/shared/loggerutils"
	"github.com/assafvayner/fs3/server/shared/tlsutils"
	"google.golang.org/grpc"
)

func main() {
	logger := loggerutils.InitLogger(config.GetName())

	ln, err := net.Listen("tcp", fmt.Sprint(":", config.GetPort()))

	if err != nil {
		logger.Fatalln(err)
		os.Exit(1)
	}
	defer ln.Close()

	tlsCredentials, err := tlsutils.GetServerTLSCredentials()
	if err != nil {
		logger.Fatalln(err)
	}

	var server *grpc.Server

	if config.IsPrimary() {
		server = grpc.NewServer(
			grpc.Creds(tlsCredentials),
			grpc.ChainUnaryInterceptor(interceptor.GetAuthInterceptor(logger)),
		)
		fs3.RegisterFs3Server(server, primary.NewPrimaryHandler(logger))
	} else {
		server = grpc.NewServer(grpc.Creds(tlsCredentials))
		primarybackup.RegisterBackupServer(server, backup.NewBackupHandler(logger))
	}

	err = server.Serve(ln)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start server serving")
		os.Exit(1)
	}
}
