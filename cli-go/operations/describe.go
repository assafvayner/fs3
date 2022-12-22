package operations

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/assafvayner/fs3/cli-go/utils"
	"github.com/assafvayner/fs3/protos/fs3"
)

func Describe(path string) {
	if path == "" {
		fmt.Fprintln(os.Stderr, "path is missing")
		os.Exit(1)
	}

	tokenString := utils.GetToken()
	req := &fs3.DescribeRequest{
		Path:  path,
		Token: tokenString,
	}

	client := utils.GetFs3Client()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := client.Describe(ctx, req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making request to server: %s\n", err)
	}

	utils.CheckFs3Status(reply.GetStatus())

	if file := reply.GetFile(); file != nil {
		fmt.Printf("file: %s\n", file.Filename)
		os.Exit(0)
	}

	dir := reply.GetDirectory()
	fmt.Println(dir.Directoryname)
	if files := dir.GetFiles(); len(files) != 0 {
		fmt.Println(" files:")
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	} else {
		fmt.Println(" no files")
	}
	if subdirs := dir.GetSubdirectories(); len(subdirs) != 0 {
		fmt.Println(" subdirectories:")
		for _, subdir := range subdirs {
			fmt.Printf("  %s\n", subdir)
		}
	} else {
		fmt.Println(" no subdirectories")
	}
}
