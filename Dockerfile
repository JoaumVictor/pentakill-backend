FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "main.go"]