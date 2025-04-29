###################################
## Build
###################################
FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

ARG TARGETOS
ARG TARGETARCH
ARG GOMODCACHE="/root/.cache/go-build"
ARG GOCACHE="/go/pkg"
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /app
COPY . .
RUN --mount=type=cache,target="$GOMODCACHE" \
    --mount=type=cache,target="$GOCACHE" \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    CGO_ENABLED=0 \
    go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o testkube-operator cmd/main.go

###################################
## Debug
###################################
FROM golang:1.23-alpine AS debug

ENV GOTRACEBACK=all
RUN go install github.com/go-delve/delve/cmd/dlv@v1.23.1

RUN apk --no-cache --update add ca-certificates && (rm -rf /var/cache/apk/* || 0)

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/testkube-operator /testkube/

ENTRYPOINT ["/go/bin/dlv", "exec", "--headless", "--continue", "--accept-multiclient", "--listen=:56401", "--api-version=2", "/testkube/testkube-operator"]

###################################
## Distribution
###################################
FROM gcr.io/distroless/static AS dist

COPY LICENSE /testkube/
COPY --from=builder /app/testkube-operator /manager

EXPOSE 8080
ENTRYPOINT ["/manager"]
