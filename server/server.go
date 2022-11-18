package server

import (
	"fmt"
  "net"
  "os"
  
	"google.golang.org/grpc"
  "gitlab.cs.washington.edu/assafv/fs3/server/fs3handler"
  "gitlab.cs.washington.edu/assafv/fs3/server/primarybackuphandler"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
  primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

func main() {
  ln, err := net.Listen("tcp", ":5000")

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
  defer ln.Close()

  server := grpc.NewServer()

  if isPrimary() {
    fs3.RegisterFs3Server(server, fs3handler.NewFs3Handler())
  } else {
    primarybackup.RegisterBackupServer(server, primarybackuphandler.NewPrimaryBackupHandler())
  }

  server.Serve(ln)
}

func isPrimary() (bool) {
  // figure out if we are primary based on config
  return true
}