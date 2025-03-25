# Gunakan base image Golang
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["/app/main"]