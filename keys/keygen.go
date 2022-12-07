package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// simple program to generate keys
// heavily inspired by https://gist.github.com/jlhawn/fd35d18ca7342a4fe000

// from keys directory run `go run keygen <optional private key file name>`

func main() {
	filename := "id_ecdsa"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	key, err := generatePrivateKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in generating private key: %s\n", err)
		os.Exit(1)
	}

	err = saveKeys(key, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error saving keys to files: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ecdsa key: %s", err)
	}
	return key, nil
}

func saveKeys(key *ecdsa.PrivateKey, filename string) error {
	privateKeyMarshalled, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return errors.New("failed to marshall private key")
	}

	privateKeyBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyMarshalled,
	}

	privateKeyFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to open file %s for writing: %s", filename, err)
	}
	defer privateKeyFile.Close()

	publicKey := &key.PublicKey
	publicKeyMarshalled, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return errors.New("failed to marshall public key")
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyMarshalled,
	}

	publicKeyFile, err := os.Create(filename + ".pub")
	if err != nil {
		return fmt.Errorf("Failed to open file %s for writing: %s", filename, err)
	}
	defer publicKeyFile.Close()

	err = pem.Encode(privateKeyFile, &privateKeyBlock)
	if err != nil {
		return fmt.Errorf("failed to write private key in pem format: %s", err)
	}
	err = pem.Encode(publicKeyFile, &publicKeyBlock)
	if err != nil {
		return fmt.Errorf("failed to write public key in pem format: %s", err)
	}
	return nil
}
