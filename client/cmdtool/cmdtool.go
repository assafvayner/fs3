package main

import (
	"context"
	"fmt"
	pb "gitlab.com/assafv/fs3/protos/fs3"
	"google.golang.org/grpc"
	"os"
)

func main() {
	// read options from command line then use grpc client
	// ran client side
	conn, err := grpc.Dial("server0:5000")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := pb.NewFs3Client(conn)

	// how to use client now:
	context := context.Background()

	request := pb.GetRequest{
		FilePath: "a/b/c",
	}

	reply, err := client.Get(context, &request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// do something with reply
	fmt.Println(reply.FilePath)
	fmt.Println(reply.Status.String())
}
