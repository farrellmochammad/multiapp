FROM golang:latest as build-env
# All these steps will be cached
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
