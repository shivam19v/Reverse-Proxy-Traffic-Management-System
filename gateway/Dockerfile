FROM golang:1.26.1-alpine

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

RUN go mod download

# Copy ALL source code
COPY . .

# Build gateway binary
RUN go build -o gateway ./cmd/gateway

EXPOSE 8080

CMD ["./gateway"]