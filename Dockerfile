ARG ALPINE_IMAGE
FROM  alpine:3.18.0

COPY ./bin/manager /manager
ENTRYPOINT ["/manager"]
