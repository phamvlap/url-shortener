ARG GOLANG_IMAGE_VERSION=1.26.8-alpine3.23

# Stage 1: Build the Go application
FROM golang:${GOLANG_IMAGE_VERSION} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o main ./cmd/server

# Stage 2: Create a minimal image for the application
FROM builder AS final

RUN addgroup -S appuser && adduser -S -G appuser appuser

WORKDIR /app

COPY --from=builder /app/main .

USER appuser

EXPOSE 8000

CMD ["./main"]
