# Gunakan base image Golang
FROM golang:1.21-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o api main.go

# Jalankan API
CMD ["/app/api"]
