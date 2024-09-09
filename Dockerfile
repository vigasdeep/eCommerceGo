# Start from the official Golang image to have a proper environment.
FROM golang:1.22@sha256:2bd56f00ff47baf33e64eae7996b65846c7cb5e0a46e0a882ef179fd89654afa as builder

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
FROM alpine:latest@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file and the .env file from the previous stage.
COPY --from=builder /app/myapp .
#COPY --from=builder /app/.env .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./myapp"]
