FROM golang:1.11.3-alpine3.8

RUN mkdir /app
ADD *.go /app/
ADD src/ /app/src/
WORKDIR /app

ARG COMMIT_REF
ARG BUILD_DATE
ENV APP_COMMIT_REF=${COMMIT_REF} \
    APP_BUILD_DATE=${BUILD_DATE}

RUN go build -o main .
EXPOSE 1234

CMD ["/app/main"]
