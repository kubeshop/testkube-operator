FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY testkube-operator /manager
ENTRYPOINT ["/manager"]
