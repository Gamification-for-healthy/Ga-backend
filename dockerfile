# Use an official Go image as a base
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the app runs on (adjust as needed)
EXPOSE 8080

# Run the application
CMD ["./main"]
