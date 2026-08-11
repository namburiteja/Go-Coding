# Build from repo root:
# docker build -f gateway/Dockerfile -t gateway:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY gateway ./gateway

WORKDIR /repo/gateway

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/gateway ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/gateway .

EXPOSE 9090

ENTRYPOINT ["./gateway"]