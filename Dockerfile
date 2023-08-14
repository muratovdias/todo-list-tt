FROM golang:1.19-alpine AS builder

WORKDIR /app
RUN export GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ .

CMD ["/app/main"]
