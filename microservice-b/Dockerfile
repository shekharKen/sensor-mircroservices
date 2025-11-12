FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service-b cmd/main.go

ENV DB_USER=root
ENV DB_PASS=root
ENV DB_HOST=mysql
ENV DB_PORT=3306
ENV DB_NAME=sensors

EXPOSE 8081 50051

CMD ["./service-b"]
