FROM golang:1.23-alpine AS builder
ARG GOOS=linux
ARG GOARCH=amd64
RUN mkdir /build
WORKDIR /build
COPY cmd/main/ ./cmd/main
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-w -s" -o server cmd/main/*

FROM scratch
COPY --from=builder /build/server /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/server"]
