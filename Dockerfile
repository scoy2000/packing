FROM golang:latest

WORKDIR /packing

COPY . .

RUN go build -o packing

EXPOSE 8080

CMD ["./packing"]