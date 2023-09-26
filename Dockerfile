ARG ALPINE_IMAGE=alpine:3.18.3
FROM ${ALPINE_IMAGE}

COPY testkube-operator /manager
ENTRYPOINT ["/manager"]
