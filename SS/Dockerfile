FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

RUN go get github.com/gorilla/mux

RUN go get github.com/gin-gonic/gin

CMD ["/app/main"]