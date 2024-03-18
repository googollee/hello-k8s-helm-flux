FROM golang:1.22.1-bookworm AS builder
COPY go.* ./src/hello/
COPY src/ ./src/hello/src/
RUN cd ./src/hello && go build -o /hello ./src

FROM busybox:1.36.1-glibc
COPY --from=builder /hello /hello
