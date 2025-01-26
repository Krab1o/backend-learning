FROM golang:latest

# Set working directory inside container
WORKDIR /app

# Copy project files
COPY . .

WORKDIR /app/cmd

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Expose application port
EXPOSE 8080

# Run Air for hot reloading
CMD ["air"]
