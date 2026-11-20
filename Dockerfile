FROM golang:1.27.1-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /ecommerce-api ./cmd

FROM alpine:3.22

RUN apk add --no-cache ca-certificates
COPY --from=builder /ecommerce-api /usr/local/bin/ecommerce-api

USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/ecommerce-api"]
