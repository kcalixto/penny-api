FROM golang:1.17.10-alpine3.16
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN goos=linux go build -o main main.go
CMD ["./main"]