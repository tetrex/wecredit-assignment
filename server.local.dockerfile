FROM golang:alpine


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 80

CMD ["air", "-c", ".air.toml"]
