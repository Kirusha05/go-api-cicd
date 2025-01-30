FROM cimg/go:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./main ./cmd/main.go

CMD ["./main"]