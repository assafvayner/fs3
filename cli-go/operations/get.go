package operations

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/assafvayner/fs3/cli-go/utils"
	fs3 "github.com/assafvayner/fs3/protos/fs3"
)

func Get(remotePath string, localFile *os.File) {
	tokenString := utils.GetToken()

	client := utils.GetFs3Client()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fs3.GetRequest{
		FilePath: remotePath,
		Token:    tokenString,
	}

	reply, err := client.Get(ctx, req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	utils.CheckFilePaths(remotePath, reply.GetFilePath())
	utils.CheckFs3Status(reply.GetStatus())

	_, err = localFile.Write(reply.GetFileContent())
	localFile.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
