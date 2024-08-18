FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o telematik-service
CMD ["./telematik-service"]
