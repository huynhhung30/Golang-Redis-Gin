FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main cmd/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -a installsuffix cgo -o main .
EXPOSE 5001


CMD ["./main"]
