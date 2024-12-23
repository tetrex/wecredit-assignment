FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./main_api ./cmd/main.go

EXPOSE 80

CMD ["./main_api"]
