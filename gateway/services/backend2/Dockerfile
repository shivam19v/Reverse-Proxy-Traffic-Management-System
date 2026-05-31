FROM golang:1.26.1-alpine

WORKDIR /app

COPY . .

RUN go build -o backend2 ./services/backend2

EXPOSE 9002

CMD ["./backend2"]