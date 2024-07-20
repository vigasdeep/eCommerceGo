# Start from the official Golang image to have a proper environment.
FROM golang:1.22 as builder

# Create a directory where our application will be placed.
WORKDIR /app

# Copy the source code and the .env file into the container.
COPY . .

# Download necessary Go modules.
RUN go mod download

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Start a new stage from scratch for a smaller, final image.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file and the .env file from the previous stage.
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./myapp"]