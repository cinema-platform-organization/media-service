FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/media-service ./cmd/main.go


FROM alpine:3.20 AS runner

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/media-service ./media-service

CMD ["./media-service"]