# original source: https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph

rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 1825 -nodes -keyout fs3-ca-key.pem -out fs3-ca-cert.pem -subj "/C=US/CN=fs3 root CA"

echo "CA's self-signed certificate"
openssl x509 -in fs3-ca-cert.pem -noout -text

for servname in primary backup auth frontend
do
  rm -rf ${servname}
  mkdir ${servname}

  # 2. Generate web server's private key and certificate signing request (CSR)
  openssl req -newkey rsa:4096 -nodes -keyout ${servname}/server-key.pem -out ${servname}/server-req.pem -subj "/C=US/CN=fs3 ${servname}"

  # 3. Use CA's private key to sign web server's CSR and get back the signed certificate
  echo "subjectAltName=DNS:${servname}.fs3\n" > ${servname}/server-ext.cnf
  openssl x509 -req -in ${servname}/server-req.pem -days 60 -CA fs3-ca-cert.pem -CAkey fs3-ca-key.pem -CAcreateserial -out ${servname}/server-cert.pem -extfile ${servname}/server-ext.cnf

  echo "Server's signed certificate"
  openssl x509 -in ${servname}/server-cert.pem -noout -text 
done
