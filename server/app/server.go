package main

import (
	"fmt"
	"log"
	"net"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/backup"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/primary"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/primary/interceptor"
	"google.golang.org/grpc"
)

func main() {
	logger := InitLogger()

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

func InitLogger() *log.Logger {
	var prefix, name string
	if config.IsPrimary() {
		name = "primary"
	} else {
		name = "backup"
	}
	prefix = "/log/"
	logFile, err := os.OpenFile(prefix+name+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open log file"+prefix+name+".log")
		os.Exit(1)
	}
	return log.New(logFile, name+": ", log.LstdFlags|log.Llongfile|log.Lmsgprefix)
}