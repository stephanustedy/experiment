FROM golang:1.11.3-alpine3.8
RUN mkdir /app
ADD *.go /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
