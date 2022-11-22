package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gitlab.cs.washington.edu/assafv/fs3/client/cmdtool/args"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// read options from command line then use grpc client
	// ran client side
	operation, filePath, err := args.ParseArgs()
	fmt.Println(operation)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, filePath+" does not exits")
		os.Exit(1)
	}

	conn, err := grpc.Dial("server0:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := fs3.NewFs3Client(conn)

	// how to use client now:
	context := context.Background()

	switch operation {
	case "get":
		request := fs3.GetRequest{
			FilePath: filePath,
		}
		reply, err := client.Get(context, &request)
		checkReply(err, reply.FilePath, reply.Status.String())
	case "cp":
		request := fs3.CopyRequest{
			FilePath: filePath,
		}
		reply, err := client.Copy(context, &request)
		checkReply(err, reply.FilePath, reply.Status.String())
	case "rm":
		request := fs3.RemoveRequest{
			FilePath: filePath,
		}
		reply, err := client.Remove(context, &request)
		checkReply(err, reply.FilePath, reply.Status.String())
	default:
		fmt.Fprintln(os.Stderr, "Operation "+operation+"does not exits")
		os.Exit(1)
	}
}

func checkReply(err error, filePath string, status string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(filePath)
	fmt.Println(status)
}
