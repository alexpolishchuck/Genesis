FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o currency_rate_service .

EXPOSE 5000

ENTRYPOINT ["/app/currency_rate_service"]