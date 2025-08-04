# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./api/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 9090
CMD ["./myapp"]