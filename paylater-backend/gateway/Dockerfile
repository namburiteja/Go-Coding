# build from repo root: docker build -f gateway/Dockerfile .
FROM golang:1.26-alpine AS build
WORKDIR /repo
COPY shared ./shared
COPY gateway ./gateway
WORKDIR /repo/gateway
RUN go mod download && CGO_ENABLED=0 go build -o /bin/gateway ./cmd
FROM alpine:3.20
COPY --from=build /bin/gateway /gateway
EXPOSE 9090
ENTRYPOINT ["/gateway"]
