# Use the official Microsoft Go dev container image
FROM mcr.microsoft.com/devcontainers/go:1-1.24-bookworm

# --- Install system dependencies ---
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        protobuf-compiler \
        git && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# --- Install Go protoc plugins ---
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Ensure Go bin is on PATH (should already be set in base image)
ENV PATH="$PATH:$(go env GOPATH)/bin"