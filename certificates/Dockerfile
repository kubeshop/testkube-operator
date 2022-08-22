FROM alpine:3.16 

RUN apk update && \
    apk add --no-cache openssl curl gettext gawk && \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin/kubectl

COPY certificates/secret.yaml secret.yaml
COPY certificates/create_cert.sh create_cert.sh
RUN chmod +x create_cert.sh
CMD ./create_cert.sh
