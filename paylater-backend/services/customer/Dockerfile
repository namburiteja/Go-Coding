# Build from repo root:
# docker build -f services/customer/Dockerfile -t customer-service:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY services/customer ./services/customer

WORKDIR /repo/services/customer

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/customer ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/customer .

EXPOSE 9093

ENTRYPOINT ["./customer"]