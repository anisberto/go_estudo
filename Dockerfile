FROM golang:1.21.5

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD go run main.go
