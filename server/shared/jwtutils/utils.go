package jwtutils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

const INVALID_USERNAME = ""
const JWT_PUBLIC_KEY_FILE_ENV_VAR = "JWT_PUBLIC_KEY_FILE"
const DEFAULT_JWT_PUBLIC_KEY_FILE = "/keys/id_ecdsa.pub"

var jwtPublicKey *ecdsa.PublicKey
var failedToRetreiveKey = false
var jwtParser *jwt.Parser = nil

type Fs3JwtClaims struct {
	Username string `json:"user"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string, validate bool) (*jwt.Token, error) {
	if jwtParser == nil {
		jwtParser = jwt.NewParser()
	}
	if !validate {
		token, _, err := jwtParser.ParseUnverified(tokenString, &Fs3JwtClaims{})
		return token, err
	}
	token, err := jwtParser.ParseWithClaims(tokenString, &Fs3JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtPublicKey()
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing token: %s\n", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Fprintln(os.Stderr, "token is invalid")
		return nil, errors.New("token in invalid")
	}
	return token, nil
}

func GetClaims(token *jwt.Token) (*Fs3JwtClaims, error) {
	if claims, ok := token.Claims.(*Fs3JwtClaims); ok {
		return claims, nil
	}
	return nil, errors.New("malformed claims")
}

func getUsernameFromToken(tokenString string, verify bool) (string, error) {
	token, err := ParseToken(tokenString, verify)
	if err != nil {
		return "", err
	}
	claims, err := GetClaims(token)
	if err != nil || claims.Valid() != nil {
		return "", err
	}
	return claims.Username, nil
}

func GetUsernameFromToken(tokenString string) (string, error) {
	return getUsernameFromToken(tokenString, true)
}

func GetUsernameFromTokenNoVerify(tokenString string) string {
	username, _ := getUsernameFromToken(tokenString, false)
	return username
}

func getJwtPublicKey() (*ecdsa.PublicKey, error) {
	if failedToRetreiveKey {
		return nil, errors.New("Failed to retrieve key")
	}
	if jwtPublicKey != nil {
		return jwtPublicKey, nil
	}
	filepath, gotten := os.LookupEnv(JWT_PUBLIC_KEY_FILE_ENV_VAR)
	if !gotten {
		filepath = DEFAULT_JWT_PUBLIC_KEY_FILE
	}
	publicKeyPemBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetJwtPublicKey: Could not read public key from %s\n", filepath)
		failedToRetreiveKey = true
		return nil, err
	}
	publicKeyPem, _ := pem.Decode(publicKeyPemBytes)
	key, err := parsePublicKey(publicKeyPem.Bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetJwtPublicKey: Failed to parse public key from %s\n", filepath)
		return nil, err
	}

	jwtPublicKey = key
	return jwtPublicKey, nil
}

func parsePublicKey(keyBytes []byte) (*ecdsa.PublicKey, error) {
	var key crypto.PublicKey
	var err error
	key, err = x509.ParsePKIXPublicKey(keyBytes)
	if err == nil {
		fmt.Fprintln(os.Stderr, "parsed public key of format PKIX")
		return key.(*ecdsa.PublicKey), nil
	}
	return nil, errors.New("Failed to parse public key with any format")
}
