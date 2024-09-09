FROM golang:1.22.6

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

ENV PORT=8080

CMD ["./main"]
