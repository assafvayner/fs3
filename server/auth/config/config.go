package config

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_AUTH_PORT = 6709
const PORT_ENV_VAR = "PORT"
const REDIS_PORT_ENV_VAR = "REDIS_PORT"
const REDIS_HOSTNAME_ENV_VAR = "REDIS_HOSTNAME"
const REDIS_PASSWORD_ENV_VAR = "REDIS_PASSWORD"
const JWT_PRIVATE_KEY_FILE_ENV_VAR = "JWT_PRIVATE_KEY_FILE"
const DEFAULT_REDIS_PORT = 6379
const DEFAULT_REDIS_HOSTNAME = "redis"
const DEFAULT_JWT_PRIVATE_KEY_FILE = "keys/id_ecdsa"

var jwtPrivateKey crypto.PrivateKey
var failedToRetreiveKey = false

func GetPort() int {
	portStr, gotten := os.LookupEnv(PORT_ENV_VAR)
	if !gotten {
		return DEFAULT_AUTH_PORT
	}
	port, err := strconv.Atoi(strings.TrimSpace(portStr))
	if err != nil {
		return DEFAULT_AUTH_PORT
	}
	return port
}

func GetRedisAddress() string {
	var port int
	var err error
	portStr, gotten := os.LookupEnv(REDIS_PORT_ENV_VAR)
	if gotten {
		port, err = strconv.Atoi(strings.TrimSpace(portStr))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to convert port %s to int", portStr)
			port = DEFAULT_REDIS_PORT
		}
	} else {
		port = DEFAULT_REDIS_PORT
	}
	hostname, gotten := os.LookupEnv(REDIS_HOSTNAME_ENV_VAR)
	if !gotten {
		hostname = DEFAULT_REDIS_HOSTNAME
	}

	return fmt.Sprintf("%s:%d", strings.TrimSpace(hostname), port)
}

func GetRedisPassword() string {
	password, gotten := os.LookupEnv(REDIS_PASSWORD_ENV_VAR)
	if !gotten {
		return ""
	}
	return strings.TrimSpace(password)
}

func GetJwtPrivateKey() (crypto.PrivateKey, error) {
	if failedToRetreiveKey {
		return nil, errors.New("failed to retreive key")
	}
	if jwtPrivateKey != nil {
		return jwtPrivateKey, nil
	}
	filepath, gotten := os.LookupEnv(JWT_PRIVATE_KEY_FILE_ENV_VAR)
	if !gotten {
		filepath = DEFAULT_JWT_PRIVATE_KEY_FILE
	}
	privateKeyPemBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetJwtPrivateKey: Could not read private key from %s\n", filepath)
		failedToRetreiveKey = true
		return nil, err
	}
	privateKeyPem, _ := pem.Decode(privateKeyPemBytes)
	key, err := parsePrivateKey(privateKeyPem.Bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetJwtPrivateKey: Failed to parse private key from %s\n", filepath)
		return nil, err
	}

	jwtPrivateKey = key
	return jwtPrivateKey, nil
}

func parsePrivateKey(keyBytes []byte) (crypto.PrivateKey, error) {
	var key crypto.PrivateKey
	var err error
	key, err = x509.ParseECPrivateKey(keyBytes)
	if err == nil {
		fmt.Fprintln(os.Stderr, "parsed private key of format EC")
		return key, nil
	}
	key, err = x509.ParsePKCS8PrivateKey(keyBytes)
	if err == nil {
		fmt.Fprintln(os.Stderr, "parsed private key of format PKCS8")
		return key, nil
	}
	key, err = x509.ParsePKCS1PrivateKey(keyBytes)
	if err == nil {
		fmt.Fprintln(os.Stderr, "parsed private key of format PKCS1")
		return key, nil
	}

	return nil, errors.New("Failed to parse private key with any format")
}
