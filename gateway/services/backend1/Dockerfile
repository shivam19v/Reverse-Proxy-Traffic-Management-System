FROM golang:1.26.1-alpine

WORKDIR /app

COPY . .

RUN go build -o backend1 ./services/backend1

EXPOSE 9001

CMD ["./backend1"]