FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ./main_api ./cmd/main.go

EXPOSE 80

CMD ["./main_api"]
