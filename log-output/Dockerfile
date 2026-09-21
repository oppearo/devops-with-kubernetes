# Stage 1: build

FROM golang:1.25.13 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app

# Stage 2: running the built binary

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=builder /todo-app /todo-app

CMD ["/todo-app"]

