# ================================
# Stage 1: Build the Go binary
# ================================

# Use official Go image based on Alpine Linux for smaller size
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all source code into the container
COPY . .

# Build the Go binary and output it as 'server'
# This compiles main.go and produces an executable suitable for Linux
RUN go build -o server main.go

# ================================
# Stage 2: Minimal runtime image
# ================================

# Use lightweight Alpine Linux as the final image
# This keeps the runtime image small and secure
FROM alpine:3.18

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Copy the 'data' folder containing products.json
COPY --from=builder /app/data ./data

# Expose port 8080 to allow external access to the Go server
EXPOSE 8080

# Define the default command to run the server binary
CMD ["./server"]
