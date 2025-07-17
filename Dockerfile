FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /payment-app

EXPOSE 8080

CMD ["/payment-app"]
