package operations

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/utils"
	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
)

func NewUser(username, password string) {
	if username == "" || password == "" {
		fmt.Fprintln(os.Stderr, "username and password may not be empty")
		os.Exit(1)
	}
	client := utils.GetAuthClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &authservice.NewUserRequest{
		Username: username,
		Password: password,
	}

	reply, err := client.NewUser(ctx, req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error from fs3 auth server: %s\n", err)
		os.Exit(1)
	}
	if !reply.GetStatus().GetSuccess() {
		fmt.Fprintln(os.Stderr, reply.GetStatus().GetMessage())
		os.Exit(1)
	}
	fmt.Println(reply.GetStatus().GetMessage())
}
