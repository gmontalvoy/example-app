FROM golang:1.18 as builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
