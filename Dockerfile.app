FROM golang:1.17 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:3.14.2
WORKDIR /app
COPY --from=builder ["/app/main", "/app"]
CMD ["/app/main"]
