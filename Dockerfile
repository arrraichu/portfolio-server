FROM golang:1.26.2-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates && \
  addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=builder /app/server .

USER app

EXPOSE 8080

ENTRYPOINT ["./server"]
