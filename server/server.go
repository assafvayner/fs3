package main

import (
	"fmt"
	"net"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
	"gitlab.cs.washington.edu/assafv/fs3/server/backup"
	"gitlab.cs.washington.edu/assafv/fs3/server/primary"
	"gitlab.cs.washington.edu/assafv/fs3/server/config"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", fmt.Sprint(":", GetPort()))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer ln.Close()

	server := grpc.NewServer()

	if config.IsPrimary() {
		fs3.RegisterFs3Server(server, primary.NewPrimaryHandler())
	} else {
		primarybackup.RegisterBackupServer(server, backup.NewBackupHandler())
	}

	server.Serve(ln)
}

func GetPort() (int) {
	if config.IsPrimary() {
		return config.PRIMARY_PORT
	}
	return config.BACKUP_PORT
}