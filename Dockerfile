# Build stage
FROM golang:1.23.1-alpine3.20 AS builder
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Install curl and download the migrate tool
RUN apk add --no-cache curl
# Download and extract migrate binary to /app
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar -xz

# Run stage
FROM alpine:3.20
WORKDIR /app

# Copy the built application and necessary files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration


# Expose the application port
EXPOSE 3000

# Use the start.sh script as the entry point
ENTRYPOINT [ "/app/start.sh" ]

# Default command to run the application
CMD [ "/app/main" ]