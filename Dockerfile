FROM golang:1.23-alpine

RUN apk add --no-cache git build-base

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN /go/bin/swag init

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080

CMD ["/main"]
