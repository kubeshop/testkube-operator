FROM gcr.io/distroless/static:nonroot
USER 65532:65532

COPY dist/manager /manager
ENTRYPOINT ["/manager"]
