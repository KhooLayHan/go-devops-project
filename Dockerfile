# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o /app/bin/go-devops-app ./cmd

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/go-devops-app .
EXPOSE 8080
CMD ["./go-devops-app"]