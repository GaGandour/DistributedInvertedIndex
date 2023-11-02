# Use an official Go runtime as the base image
FROM golang:1.20

# Set the working directory in the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o dii .

# Run the application
CMD ["./dii"]
