package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/service"
)

func main() {
	logger := InitLogger()

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

func InitLogger() *log.Logger {
	var prefix, name string
	prefix = "/log/"
	logFile, err := os.OpenFile(prefix+"authservice.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open log file")
		os.Exit(1)
	}
	return log.New(logFile, name+": ", log.LstdFlags|log.Llongfile|log.Lmsgprefix)
}
