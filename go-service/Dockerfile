FROM golang:latest as build-env
# All these steps will be cached
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build

CMD ["./go-service"]
