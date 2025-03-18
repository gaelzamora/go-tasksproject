FROM golang:1.23 AS builder

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM golang:1.23 AS runner

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /go/bin/CompileDaemon /go/bin/CompileDaemon

EXPOSE 3000

CMD ["/go/bin/CompileDaemon", "--build=go build -o main ./cmd/main.go", "--command=./main"]
