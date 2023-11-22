FROM golang:1.21.3-alpine3.18 as builder

ARG SRC_DIR=/go/towerinfo
RUN mkdir -p $SRC_DIR

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPATH=/go

WORKDIR $SRC_DIR

COPY . .

RUN go build -a -o towerinfo-api ./services/tower-profile-api


FROM alpine:3.18.3

RUN addgroup -g 1000 -S api-user && \
    adduser -u 1000 -S api-user -G api-user && \
    mkdir /towerinfo-api && \
    chown 1000:1000 /towerinfo-api

ARG SRC_DIR=/go/towerinfo

COPY --from=builder ${SRC_DIR}/towerinfo-api /usr/bin/towerinfo-api

USER 1000

EXPOSE 4001
ENTRYPOINT ["/usr/bin/towerinfo-api"]



