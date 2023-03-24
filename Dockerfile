FROM alpine
COPY testkube-operator /manager
ENTRYPOINT ["/manager"]
