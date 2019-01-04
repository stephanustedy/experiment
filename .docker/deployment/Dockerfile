FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://storage.googleapis.com/golang/go1.11.3.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.11.3.linux-amd64.tar.gz
RUN rm /tmp/go1.11.3.linux-amd64.tar.gz

RUN mkdir -p /app/src/github.com/stephanustedy/experiment
ENV GOPATH /app
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV WORKDIR /app/src/github.com/stephanustedy/experiment

ADD *.go $WORKDIR
ADD src/ $WORKDIR/src
ADD files/ /

WORKDIR $WORKDIR

ARG COMMIT_REF
ARG BUILD_DATE
ENV APP_COMMIT_REF=${COMMIT_REF} \
    APP_BUILD_DATE=${BUILD_DATE}

RUN go get
RUN go build -o main .

EXPOSE 1234

CMD ["/app/src/github.com/stephanustedy/experiment/main"]
