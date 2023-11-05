# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /gateway
WORKDIR /gateway

# Copy the API gateway code into the container
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application

RUN CGO_ENABLED=0 GOOS=linux go build -o ./gateway

RUN chmod +x ./gateway


# Expose port 8080
EXPOSE 8080

# Run the API gateway executable
CMD ["./gateway"]
