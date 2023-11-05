# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /service2
WORKDIR /service2

# Copy the service2 code into the container
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application

RUN CGO_ENABLED=0 GOOS=linux go build -o ./service2

RUN chmod +x ./service2

# Expose port 8081
EXPOSE 8082

# Run the service1 executable
CMD ["./service2"]
