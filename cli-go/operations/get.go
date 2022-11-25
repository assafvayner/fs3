package operations

import (
  "context"
  "fmt"
  "os"
  "time"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/utils"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func Get(remotePath string, localFile *os.File) {
	client := utils.GetFs3Client()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fs3.GetRequest{
		FilePath:    remotePath,
	}

	reply, err := client.Get(ctx, req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

  utils.CheckFilePaths(remotePath, reply.GetFilePath())
  utils.CheckStatus(reply.GetStatus())

  err = os.WriteFile(localFile.Name(), reply.GetFileContent(), 0666)
  localFile.Close()
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
