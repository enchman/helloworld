FROM golang:1.22-alpine as builder

WORKDIR /app

# Download Go modules
COPY go.* ./
RUN go mod download

COPY . ./

# Build
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -v -o server

# Prepare container
FROM alpine:latest

LABEL author="EnchMan"

COPY --from=builder /app/server /app/server

# Run
CMD ["/app/server", "-port", "80"]

EXPOSE 80
