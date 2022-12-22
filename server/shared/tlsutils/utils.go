package tlsutils

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

const CERT_PATH_ENV_VAR = "FS3_SIGNED_CERTIFICATE_PATH"
const KEY_PATH_ENV_VAR = "FS3_SIGNED_KEY_PATH"
const CA_CERT_PATH_ENV_VAR = "FS3_CA_PATH"
const DEFAULT_CERT_PATH = "/certificates/server-cert.pem"
const DEFAULT_KEY_PATH = "/certifcates/server-key.pem"
const DEFAULT_CA_CERT_PATH = "/certificates/fs3-ca-cert.pem"

var serverTlsCredentials credentials.TransportCredentials = nil
var clientTlsCredentials credentials.TransportCredentials = nil

func GetServerTLSCredentials() (credentials.TransportCredentials, error) {
	if serverTlsCredentials != nil {
		return serverTlsCredentials, nil
	}
	// https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph
	certPath := DEFAULT_CERT_PATH
	keyPath := DEFAULT_KEY_PATH
	if certPathConfigged, gotten := os.LookupEnv(CERT_PATH_ENV_VAR); gotten {
		certPath = certPathConfigged
	}
	if keyPathConfigged, gotten := os.LookupEnv(KEY_PATH_ENV_VAR); gotten {
		keyPath = keyPathConfigged
	}

	caCertPath := DEFAULT_CA_CERT_PATH
	if caCertPathConfigged, gotten := os.LookupEnv(CA_CERT_PATH_ENV_VAR); gotten {
		caCertPath = caCertPathConfigged
	}

	caCertBytes, err := os.ReadFile(caCertPath)
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCertBytes) {
		panic(errors.New("failed to append cert"))
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	serverTlsCredentials = credentials.NewTLS(config)
	return serverTlsCredentials, nil
}

func GetClientTLSCredentials() (credentials.TransportCredentials, error) {
	if clientTlsCredentials != nil {
		return clientTlsCredentials, nil
	}
	// Load certificate of the CA who signed server's certificate
	caCertPath := DEFAULT_CA_CERT_PATH
	if caCertPathConfigged, gotten := os.LookupEnv(CA_CERT_PATH_ENV_VAR); gotten {
		caCertPath = caCertPathConfigged
	}
	pemServerCA, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	certPath := DEFAULT_CERT_PATH
	keyPath := DEFAULT_KEY_PATH
	if certPathConfigged, gotten := os.LookupEnv(CERT_PATH_ENV_VAR); gotten {
		certPath = certPathConfigged
	}
	if keyPathConfigged, gotten := os.LookupEnv(KEY_PATH_ENV_VAR); gotten {
		keyPath = keyPathConfigged
	}
	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	clientTlsCredentials = credentials.NewTLS(config)
	return clientTlsCredentials, nil
}
