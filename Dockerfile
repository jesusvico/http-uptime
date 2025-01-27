ARG GO_VERSION="1.23.5"
ARG ALPINE_VERSION="3.20"

# Stagr 1: Build the binary
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /go/src/github.com/jesusvico/http-uptime
COPY . .

RUN go build -o /go/bin/http-uptime ./cmd/http-uptime

# Stage 2: Create a minimal image
FROM alpine:${ALPINE_VERSION}

COPY scripts/docker/entrypoint.sh /entrypoint.sh
COPY --from=builder /go/bin/http-uptime /usr/bin/http-uptime

ENV PORT=8080
ENV CONFIG_PATH="/etc/http-uptime/config.yaml"

CMD ["sh", "entrypoint.sh"]
