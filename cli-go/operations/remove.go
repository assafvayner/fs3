package operations

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/utils"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func Remove(remotePath string) {
	tokenString := utils.GetToken()

	client := utils.GetFs3Client()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fs3.RemoveRequest{
		FilePath: remotePath,
		Token:    tokenString,
	}

	reply, err := client.Remove(ctx, req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	utils.CheckFilePaths(remotePath, reply.GetFilePath())
	utils.CheckStatus(reply.GetStatus())
}
