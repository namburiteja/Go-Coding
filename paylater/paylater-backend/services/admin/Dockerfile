# Build from repo root:
# docker build -f services/admin/Dockerfile -t admin-service:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY services/admin ./services/admin

WORKDIR /repo/services/admin

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/admin ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/admin .

EXPOSE 9091

ENTRYPOINT ["./admin"]