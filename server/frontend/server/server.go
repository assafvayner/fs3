package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const AUTH_HOSTNAME_ENV_VAR = "AUTH_HOSTNAME"
const AUTH_PORT_ENV_VAR = "AUTH_PORT"
const PRIMARY_HOSTNAME_ENV_VAR = "PRIMARY_HOSTNAME"
const PRIMARY_PORT_ENV_VAR = "PRIMARY_PORT"

type FrontendServer struct {
	Logger     *log.Logger
	Port       int
	Fs3Client  fs3.Fs3Client
	AuthClient authservice.AuthClient
}

func NewFrontendServer(port int, logger *log.Logger) *FrontendServer {
	return &FrontendServer{
		Logger: logger,
		Port:   port,
	}
}

func (server *FrontendServer) VerifyFs3Client() {
	if server.Fs3Client != nil {
		return
	}
	conn, err := grpc.Dial(getPrimaryAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not create a connection to the primary server")
		fmt.Fprintln(os.Stderr, "guess I'll die")
		os.Exit(1)
	}
	server.Fs3Client = fs3.NewFs3Client(conn)
}

func (server *FrontendServer) VerifyAuthClient() {
	if server.Fs3Client != nil {
		return
	}
	conn, err := grpc.Dial(getAuthAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not create a connection to the auth server")
		fmt.Fprintln(os.Stderr, "guess I'll die")
		os.Exit(1)
	}
	server.AuthClient = authservice.NewAuthClient(conn)
}

func getPrimaryAddress() string {
	primaryHostname, gotten := os.LookupEnv(PRIMARY_HOSTNAME_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "could not get primary's hostname")
		return ""
	}
	primaryPortStr, gotten := os.LookupEnv(PRIMARY_PORT_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "could not get primary's port")
		return ""
	}
	primaryPort, err := strconv.Atoi(primaryPortStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "primary port is invalid format: %s\n", err)
		return ""
	}
	return fmt.Sprintf("%s:%d", primaryHostname, primaryPort)
}

func getAuthAddress() string {
	authHostname, gotten := os.LookupEnv(AUTH_HOSTNAME_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "could not get primary's hostname")
		return ""
	}
	authPortStr, gotten := os.LookupEnv(AUTH_PORT_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "could not get primary's port")
		return ""
	}
	authPort, err := strconv.Atoi(authPortStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "auth port is invalid format: %s\n", err)
		return ""
	}
	return fmt.Sprintf("%s:%d", authHostname, authPort)
}
