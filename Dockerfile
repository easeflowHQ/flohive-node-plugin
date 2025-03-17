FROM golang:1.24.1-alpine AS builder
WORKDIR /app
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    musl-dev \
    linux-headers \
    binutils-gold \
    gcompat

COPY go.mod go.sum ./
RUN go mod download

COPY plugin/ plugins/

RUN mkdir build
RUN go build -buildmode=plugin -o build/plugins/signature-scan-demo-v2.so plugins/*.go 

FROM easeflow/mawari-node:latest
WORKDIR /app
COPY --from=builder /app/build /app/
COPY config.yml /app/

ENTRYPOINT ["/app/node"]