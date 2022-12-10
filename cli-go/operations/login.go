package operations

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
	"time"

	"gitlab.cs.washington.edu/assafv/fs3/cli-go/utils"
	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"golang.org/x/term"
)

func Login(username string, useToken bool) {
	var err error
	client := utils.GetAuthClient()

	if username == "" {
		username = requestUsername()
	}

	password := ""
	tokenString := ""
	if useToken {
		tokenString, err = utils.ReadToken()
		if err != nil {
			fmt.Fprintf(os.Stderr, "token error: %s\n", err)
			tokenString = ""
		}
	}
	if !useToken || tokenString == "" {
		password = requestPassword()
	}

	request := &authservice.GetNewTokenRequest{
		Username:      username,
		Password:      password,
		PreviousToken: tokenString,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := client.GetToken(ctx, request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "login failed: %s\n", err)
	}

	if !resp.GetStatus().GetSuccess() {
		fmt.Fprintln(os.Stderr, resp.GetStatus().GetMessage())
		os.Exit(1)
	}
	err = utils.WriteToken(resp.GetToken())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write token: %s\n", err)
	}
}

func requestUsername() string {
	var res string
	fmt.Print("enter username: ")
	fmt.Scanln(&res)
	res = strings.TrimSpace(res)
	isValid := validateUsername(res)
	isEmpty := len(res) == 0
	for !isValid || isEmpty {
		if !isValid {
			fmt.Fprintln(os.Stderr, "username may not contain whitespace or special characters")
		}
		if isEmpty {
			fmt.Fprintln(os.Stderr, "username cannot be empty")
		}
		fmt.Print("retry username: ")
		fmt.Scanln(&res)
		res = strings.TrimSpace(res)
		isValid = validateUsername(res)
		isEmpty = len(res) == 0
	}
	return res
}

func requestPassword() string {
	fmt.Print("enter password: ")
	strBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read input password")
		os.Exit(1)
	}
	res := strings.TrimSpace(string(strBytes))
	isValid := validatePassword(res)
	isEmpty := len(res) == 0
	for !isValid || isEmpty {
		if !isValid {
			fmt.Fprintln(os.Stderr, "password may not contain whitespace")
		}
		if isEmpty {
			fmt.Fprintln(os.Stderr, "\npassword cannot be empty")
		}
		fmt.Print("retry password: ")
		strBytes, err = term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to read input password")
			os.Exit(1)
		}
		res = strings.TrimSpace(string(strBytes))
		isValid = validatePassword(res)
		isEmpty = len(res) == 0
	}
	fmt.Println() // fix newline alignment on terminal
	return res
}

func validateUsername(username string) bool {
	return regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`).MatchString(username)
}

func validatePassword(password string) bool {
	return !strings.ContainsAny(password, " \t\n")
}
