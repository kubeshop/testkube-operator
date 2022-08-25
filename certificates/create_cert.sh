#!/bin/sh

openssl req -x509 -nodes -days 90 \
    -subj  "/C=CA/ST=QC/O=Kubeshop /CN=$WEBHOOK_DOMAIN_NAME" \
    -addext "subjectAltName = DNS:$WEBHOOK_DOMAIN_NAME, DNS:$DNS_DOMAIN_NAME1, DNS:$DNS_DOMAIN_NAME2" \
     -newkey rsa:2048 -keyout webhook-serv.key \
     -out webhook-serv.crt;

echo WEBHOOK_DOMAIN_NAME=$WEBHOOK_DOMAIN_NAME DNS_DOMAIN_NAME1=$DNS_DOMAIN_NAME1 DNS_DOMAIN_NAME2=$DNS_DOMAIN_NAME2 NAMESPACE=$NAMESPACE
export WEBHOOK_CERT=$(base64 webhook-serv.crt | awk 'BEGIN{ORS="";} {print}')
export WEBHOOK_KEY=$(base64 webhook-serv.key | awk 'BEGIN{ORS="";} {print}')
cat secret.yaml | envsubst > /tmp/foo.yaml
cat /tmp/foo.yaml
kubectl apply -f /tmp/foo.yaml