FROM golang:1.26.1-alpine

WORKDIR /app

COPY . .

RUN go build -o backend3 ./services/backend3

EXPOSE 9003

CMD ["./backend3"]