FROM golang
RUN mkdir /app
ADD . /app
RUN go get github.com/go-sql-driver/mysql
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8082