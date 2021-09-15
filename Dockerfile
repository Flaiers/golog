FROM golang:latest

WORKDIR /usr/projects/go-logging
COPY . /usr/projects/go-logging/

RUN go get github.com/joho/godotenv github.com/lib/pq github.com/gorilla/mux
RUN go mod download