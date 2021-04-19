FROM golang:1.16

Run mkdir /app
Add . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
CMD ["/app/main"]