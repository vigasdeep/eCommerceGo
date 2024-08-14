# Start from the official Golang image to have a proper environment.
FROM golang:1.23@sha256:8e529b64d382182bb84f201dea3c72118f6ae9bc01d27190ffc5a54acf0091d2 as builder

# Create a directory where our application will be placed.
WORKDIR /app
LABEL org.opencontainers.image.source="https://github.com/vigasdeep/eCommerceGo"

# Copy the source code and the .env file into the container.
COPY . .

# Download necessary Go modules.
RUN go mod download

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Start a new stage from scratch for a smaller, final image.
FROM alpine:latest@sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file and the .env file from the previous stage.
COPY --from=builder /app/myapp .
#COPY --from=builder /app/.env .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./myapp"]
