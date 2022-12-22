package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/assafvayner/fs3/protos/authservice"
	fs3 "github.com/assafvayner/fs3/protos/fs3"
)

type Fs3JwtClaims struct {
	Username string `json:"user"`
	jwt.RegisteredClaims
}

func getServerHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "localhost"
	}
	if strings.Contains(hostname, "cloudlab") {
		return "server0"
	}
	return "localhost"
}

/*
 * creates grpc client to make call to fs3 server
 * exits if fails to make a connection
 */
func GetFs3Client() fs3.Fs3Client {
	addr := fmt.Sprintf("%s:5000", getServerHostname())
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := fs3.NewFs3Client(conn)
	return client
}

func CheckFs3Status(received fs3.Status) {
	if received == fs3.Status_GREAT_SUCCESS {
		return
	}
	fmt.Fprintln(os.Stderr, errors.New("Server returned status: "+received.String()))
	os.Exit(1)
}

func CheckFilePaths(expected, received string) {
	if received == expected {
		return
	}
	fmt.Fprintln(os.Stderr, errors.New("Server returned bad path: "+received))
	os.Exit(1)
}

func GetAuthClient() authservice.AuthClient {
	addr := fmt.Sprintf("%s:6709", getServerHostname())
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := authservice.NewAuthClient(conn)
	return client
}

func WriteToken(token string) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(homedir+"/.fs3", 0777)
	if err != nil {
		return nil
	}
	return os.WriteFile(homedir+"/.fs3/auth_token", []byte(token), 0777)
}

func ReadToken() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	bytesToken, err := os.ReadFile(homedir + "/.fs3/auth_token")
	return string(bytesToken), err
}

func GetToken() string {
	tokenString, err := ReadToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting token, you may need to log in again: %s\n", err)
		os.Exit(1)
	}
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &Fs3JwtClaims{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse token, you may need to log in again: %s\n", err)
		os.Exit(1)
	}
	claims, ok := token.Claims.(*Fs3JwtClaims)
	if !ok {
		fmt.Fprintln(os.Stderr, "failed to decode get token fields, you may need to log in again")
		os.Exit(1)
	}
	if err = claims.Valid(); err != nil {
		fmt.Fprintf(os.Stderr, "token validation failed, potentially expired, you may need to log in again: %s\n", err)
		os.Exit(1)
	}
	return tokenString
}
