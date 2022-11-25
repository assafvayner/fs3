package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

/*
 * creates grpc client to make call to fs3 server
 * exits if fails to make a connection
 */
func GetFs3Client() fs3.Fs3Client {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := fs3.NewFs3Client(conn)
	return client
}

func CheckStatus(received fs3.Status) {
  if received == fs3.Status_GREAT_SUCCESS {
    return
	}
  fmt.Fprintln(os.Stderr, errors.New("Server returned status: " + received.String()))
  os.Exit(1)
}

func CheckFilePaths(expected, received string) {
  if received == expected {
    return
  }
  fmt.Fprintln(os.Stderr, errors.New("Server returned bad path: " + received))
  os.Exit(1)
}
