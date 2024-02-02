# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
