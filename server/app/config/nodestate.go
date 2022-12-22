package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DEFAULT_PRIMARY_PORT = 5000
var DEFAULT_BACKUP_PORT = 50000
var ROLE_ENV_VAR = "ROLE"
var PORT_ENV_VAR = "PORT"
var BACKUP_PORT_ENV_VAR = "BACKUP_PORT"

var ROLE = ""

func IsPrimary() bool {
	if ROLE != "" {
		return ROLE == "primary"
	}
	role, gotten := os.LookupEnv(ROLE_ENV_VAR)
	if !gotten {
		fmt.Fprintln(os.Stderr, "No ROLE environment variable")
		os.Exit(1)
	}
	role = strings.TrimSpace(role)
	ROLE = role
	if role == "primary" {
		return true
	}
	if role != "backup" {
		fmt.Fprint(os.Stderr, role)
		fmt.Fprint(os.Stderr, "backup")
		fmt.Fprintf(
			os.Stderr,
			"Invalid ROLE environment variable, needs primary/backup got: %s\n",
			role,
		)
		os.Exit(1)
	}
	return false
}

func GetName() string {
	if IsPrimary() {
		return "primary"
	}
	return "backup"
}

func GetPort() int {
	portStr, gotten := os.LookupEnv(PORT_ENV_VAR)
	if !gotten {
		return getDefaultPort()
	}
	port, err := strconv.Atoi(strings.TrimSpace(portStr))
	if err != nil {
		return getDefaultPort()
	}
	return port
}

func getDefaultPort() int {
	if IsPrimary() {
		return DEFAULT_PRIMARY_PORT
	}
	return DEFAULT_BACKUP_PORT
}

func GetBackupPort() int {
	portStr, gotten := os.LookupEnv(BACKUP_PORT_ENV_VAR)
	if !gotten {
		return DEFAULT_BACKUP_PORT
	}
	port, err := strconv.Atoi(strings.TrimSpace(portStr))
	if err != nil {
		return DEFAULT_BACKUP_PORT
	}
	return port
}
