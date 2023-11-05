FROM golang:1.21.3-alpine3.18 as builder

ARG SRC_DIR=/go/plantinfo
RUN mkdir -p $SRC_DIR

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPATH=/go

WORKDIR $SRC_DIR

COPY . .

RUN go build -a -o plantinfo-web ./cmd/web


FROM alpine:3.18.3

RUN addgroup -g 1000 -S web-user && \
    adduser -u 1000 -S web-user -G web-user && \
    mkdir /plantinfo && \
    chown 1000:1000 /plantinfo

ARG SRC_DIR=/go/plantinfo

COPY --from=builder ${SRC_DIR}/plantinfo-web /plantinfo/plantinfo-web
RUN chown 1000:1000 /plantinfo/plantinfo-web

COPY --from=builder ${SRC_DIR}/ui /plantinfo/ui
RUN chown 1000:1000 -R /plantinfo/ui

USER 1000
EXPOSE 8090

WORKDIR "/plantinfo"

ENTRYPOINT ["./plantinfo-web"]



