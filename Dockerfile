FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api main.go

FROM alpine:latest

COPY --from=builder /app/api /app/
COPY --from=builder /app/.env /app/

WORKDIR /app

EXPOSE 8080

CMD ["./api"]