# Use official Golang image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Build the Go application
RUN go build -o url-shortener-backend cmd/main.go

# Expose port
EXPOSE 9090

# Start application
CMD ["/app/url-shortener-backend"]