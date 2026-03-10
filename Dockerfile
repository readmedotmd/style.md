# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /src

# Cache dependency downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build a fully-static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o /out/generate-server ./cmd/generate-server

# ---- Runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot

LABEL org.opencontainers.image.source="https://github.com/readmedotmd/style.md" \
      org.opencontainers.image.title="style.md generate-server" \
      org.opencontainers.image.description="Standalone HTTP server for style.md SVG badge generation" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.vendor="readmedotmd"

COPY --from=builder /out/generate-server /usr/local/bin/generate-server

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["generate-server"]
