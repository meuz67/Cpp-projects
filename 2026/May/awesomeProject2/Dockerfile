FROM golang:1.26.3-bookworm
WORKDIR /app
COPY go.sum ./
COPY go.mod ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o main .
CMD ["./main"]