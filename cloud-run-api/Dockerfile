# syntax=docker/dockerfile:1.4

# --- Builder Stage ---
# Define the base image for the builder stage
# We need an image with Go and apk (Alpine package manager)
# Using the official golang image based on Alpine is a good choice
FROM golang:1.24-alpine AS builder

# Install build dependencies (git needed for private modules if GOPRIVATE is set)
# ca-certificates might be needed for HTTPS calls during build (e.g., go mod download)
RUN apk add --no-cache git ca-certificates

# Set the working directory *within the builder stage*
WORKDIR /app

# --- Optimized Dependency Caching ---
# Copy only the files required for dependency resolution first
COPY go.mod go.sum ./

# Download dependencies using BuildKit cache mounts for efficiency
# This prevents re-downloading everything unless go.mod/go.sum changes.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download -x # -x for verbose output if needed for debugging
# Optional: Verify checksums (good practice)
# The cache mount needs to be repeated here as well
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod verify

# --- Copy Application Source Code ---
# Copy the entire source code needed for the build.
# This includes cmd/, internal/, etc. relative to the WORKDIR (/app)
COPY . .

# --- Build the Binary ---
# Build the main application binary statically.
# Using -trimpath reduces binary size by removing local paths.
# Using -ldflags="-w -s" strips debug info and symbol table, further reducing size.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-w -s" \
    -o /server ./cmd # Build the main package located in ./cmd

# --- Final Stage ---
# Use a minimal distroless static image as the final base.
# This contains only the application and its minimal runtime dependencies.
FROM gcr.io/distroless/static-debian12 AS final

# Copy *only* the built binary from the builder stage.
# The source path is absolute within the builder stage.
COPY --from=builder /server /server

# Expose the port the application listens on (matches default in main.go)
EXPOSE 8080

# Run as a non-root user for security.
# UID 65532 is the standard 'nonroot' user in distroless images.
USER 65532:65532

# Define the entrypoint for the container. This runs the application.
ENTRYPOINT ["/server"]