#!/bin/bash

# Define the namespace where the operator is running.
NAMESPACE="testkube"

# Define the names of your certificate and key files.
OUTPUT_DIR="/tmp/serving-certs"
mkdir -p ${OUTPUT_DIR}
CA_KEY="$OUTPUT_DIR/ca.key"
CA_CERT="$OUTPUT_DIR/ca.crt"
CSR_FILE="$OUTPUT_DIR/server.csr"
CERT_FILE="$OUTPUT_DIR/cert.crt"
KEY_FILE="$OUTPUT_DIR/key.key"

## Clean existing keys
#rm -f ${CERT_FILE} ${KEY_FILE} ${CA_KEY} ${CA_CERT}
#
## Create the CA cert and private key
#openssl genrsa -out ${CA_KEY} 2048
#openssl req -new -x509 -days 365 -key ${CA_KEY} -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ${CA_CERT}
#
## Create the private key, certificate signing request (CSR), and certificate
#openssl req -newkey rsa:2048 -nodes -keyout ${KEY_FILE} -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=testkube-operator-webhook-service.${NAMESPACE}.svc" -out ${CSR_FILE}
#openssl x509 -req -extfile <(printf "subjectAltName=DNS:testkube-operator-webhook-service.%s.svc" $NAMESPACE) -days 365 -in ${CSR_FILE} -CA ${CA_CERT} -CAkey ${CA_KEY} -CAcreateserial -out ${CERT_FILE}
#
## Check if the certificate and key were generated successfully.
#if [[ $? -ne 0 ]]; then
#    echo "Failed to generate the certificate and key."
#    exit 1
#fi

certs=$(self-signed-cert --namespace testkube --service-name testkube-operator-webhook-service)

# Re-create the Kubernetes secret for the webhook server cert.
kubectl delete secret webhook-server-cert -n ${NAMESPACE} --ignore-not-found=true
kubectl create secret tls webhook-server-cert \
    --cert="$certs/server.crt" \
    --key="$certs/server.key"  \
    -n ${NAMESPACE}

# Check if the secret was created successfully.
if [[ $? -eq 0 ]]; then
    echo "Secret 'webhook-server-cert' created successfully."
else
    echo "Failed to create secret 'webhook-server-cert'."
fi