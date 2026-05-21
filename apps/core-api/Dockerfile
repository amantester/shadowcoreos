FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY apps/core-api/go.mod apps/core-api/go.sum ./
RUN go mod download

COPY apps/core-api ./

RUN go build -o shadowcore-api ./cmd/server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/shadowcore-api .

EXPOSE 8080

CMD ["./shadowcore-api"]
