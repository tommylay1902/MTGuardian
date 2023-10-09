# Use the official Golang image as the base image
FROM golang:1.21.1


WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /main

EXPOSE 8080

CMD [ "/main"]