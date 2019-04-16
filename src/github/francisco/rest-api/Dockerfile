FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

ENV GOPATH=/app

RUN go build

EXPOSE 8080

CMD ["./app"]
