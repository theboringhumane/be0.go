# Use the official Go image as a base
FROM golang:1.24-rc-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o build/posthoot cmd/main.go

# Build helper binary
RUN CGO_ENABLED=0 GOOS=linux go build -o build/helper cmd/helper/main.go

# Use a minimal alpine image for the final stage
FROM gcr.io/distroless/static-debian12:nonroot

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --chmod=755 --from=builder /app/build/posthoot .
COPY --chmod=755 --from=builder /app/build/helper .

# Copy template seeder data for Airley templates
# Source: /app/internal/models/seeder/airley/templates.json
# Destination: /app/internal/models/seeder/airley/templates.json
COPY --chmod=755 --from=builder /app/internal/models/seeder/airley/templates.json /app/internal/models/seeder/airley/

# Copy all initial setup seeder files for database initialization 
# Source: /app/internal/models/seeder/initial-setup/*
# Destination: /app/internal/models/seeder/initial-setup/
COPY --chmod=755 --from=builder /app/internal/models/seeder/initial-setup/* /app/internal/models/seeder/initial-setup/

# Expose ports
EXPOSE 9001

# Set the entry point
CMD ["/app/be0"]
