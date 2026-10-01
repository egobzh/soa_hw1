# -- build stage --
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o notify .

# -- run stage --
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/notify .

EXPOSE 30000

CMD ["./notify"]
