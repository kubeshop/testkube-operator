FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY dist/manager /manager
ENTRYPOINT ["/manager"]
