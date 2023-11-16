FROM golang:1.21.3-alpine3.18 as builder

ARG SRC_DIR=/go/plantinfo
RUN mkdir -p $SRC_DIR

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPATH=/go

WORKDIR $SRC_DIR

COPY . .

RUN go build -a -o plantinfo-api ./services/plant-profile-api


FROM alpine:3.18.3

RUN addgroup -g 1000 -S api-user && \
    adduser -u 1000 -S api-user -G api-user && \
    mkdir /plantinfo-api && \
    chown 1000:1000 /plantinfo-api

ARG SRC_DIR=/go/plantinfo

COPY --from=builder ${SRC_DIR}/plantinfo-api /usr/bin/plantinfo-api

USER 1000

EXPOSE 4000
ENTRYPOINT ["/usr/bin/plantinfo-api"]



