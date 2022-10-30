FROM golang:1.19-alpine3.16

ADD . /application
WORKDIR /application

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]