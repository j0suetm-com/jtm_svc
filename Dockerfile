# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./
COPY api/ ./api
COPY util/ ./util
COPY config.json ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /jtm_svc

CMD ["/jtm_svc"]
