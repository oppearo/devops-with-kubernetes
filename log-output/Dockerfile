# Stage 1: build
FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/generator ./cmd/generator
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/contentserver ./cmd/contentserver

# Stage 2: runtime
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=builder /bin/generator /bin/generator
COPY --from=builder /bin/contentserver /bin/contentserver

# Default to running the content server; override with /bin/generator for the writer container
CMD ["/bin/contentserver"]
