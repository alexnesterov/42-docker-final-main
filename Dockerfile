FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /my_app ./cmd/main

FROM scratch
COPY --from=builder /my_app /my_app

CMD ["/my_app"]
