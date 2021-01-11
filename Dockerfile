FROM golang:latest

ADD . /go/src/temp

RUN rm -rf /go/src/temp/static

RUN ls /go/src/temp

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gin-gonic/gin"
RUN go get .
RUN go build .

EXPOSE 8080

CMD ./main
