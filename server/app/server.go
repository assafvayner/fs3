package main

import (
	"fmt"
	"net"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/backup"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/primary"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/primary/interceptor"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/loggerutils"
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

	var server *grpc.Server

	if config.IsPrimary() {
		server = grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.GetAuthInterceptor(logger)))
		fs3.RegisterFs3Server(server, primary.NewPrimaryHandler(logger))
	} else {
		server = grpc.NewServer()
		primarybackup.RegisterBackupServer(server, backup.NewBackupHandler(logger))
	}

	err = server.Serve(ln)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start server serving")
		os.Exit(1)
	}
}
