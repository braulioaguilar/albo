FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o character .

EXPOSE 9090

CMD ["./character"]
