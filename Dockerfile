FROM golang:1.18

WORKDIR /go/src

COPY . .

RUN GOOS=linux go build main.go

EXPOSE 8089

ENTRYPOINT ["./main"]
