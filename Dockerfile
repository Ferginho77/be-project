# ==========================
# Stage 1: Build Golang App
# ==========================
FROM golang:1.26 AS builder

WORKDIR /app

# Copy dependency file terlebih dahulu
COPY go.mod go.sum ./

RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ==========================
# Stage 2: Runtime
# ==========================
FROM alpine:latest

WORKDIR /root/

# Install certificate untuk HTTPS
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]