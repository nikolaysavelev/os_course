FROM golang:1.20-alpine
WORKDIR /app
COPY client.go .
RUN go build -o client client.go
CMD ["/app/client"]