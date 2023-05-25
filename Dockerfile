ARG ALPINE_IMAGE
FROM ${ALPINE_IMAGE}

COPY testkube-operator /manager
ENTRYPOINT ["/manager"]
