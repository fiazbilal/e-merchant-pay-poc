# Build Stage
FROM golang:1.22.4 AS BuildStage

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the entire project into the container
COPY . .

# Build the application binary with CGO disabled for Alpine compatibility
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payment-processor ./cmd

# Runtime Stage
FROM alpine:latest

# Add CA certificates and tzdata
RUN apk --no-cache add ca-certificates tzdata

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary and the .env file from the build stage
COPY --from=BuildStage /app/payment-processor /app/payment-processor
COPY --from=BuildStage /app/cmd/.env /app/.env

# Make sure the binary is executable
RUN chmod +x /app/payment-processor

# Expose the port the app listens on
EXPOSE 4100

# Set the entrypoint to run the application
ENTRYPOINT ["/app/payment-processor"]