FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sensor cmd/main.go

ENV SENSOR_TYPE=TEMPERATURE
ENV FREQUENCY=2

EXPOSE 8080

CMD ["./sensor"]
