package operations

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/utils"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func Copy(localFile *os.File, remotePath string) {
	tokenString := utils.GetToken()

	fileContents, err := os.ReadFile(localFile.Name())
	localFile.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := utils.GetFs3Client()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fs3.CopyRequest{
		FilePath:    remotePath,
		FileContent: fileContents,
		Token:       tokenString,
	}

	reply, err := client.Copy(ctx, req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	utils.CheckFilePaths(remotePath, reply.GetFilePath())
	utils.CheckStatus(reply.GetStatus())
}
