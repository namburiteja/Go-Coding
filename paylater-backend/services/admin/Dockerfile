# build from repo root: docker build -f services/admin/Dockerfile .
FROM golang:1.26-alpine AS build
WORKDIR /repo
COPY shared ./shared
COPY services/admin ./services/admin
WORKDIR /repo/services/admin
RUN go mod download && CGO_ENABLED=0 go build -o /bin/admin ./cmd
FROM alpine:3.20
COPY --from=build /bin/admin /admin
EXPOSE 9091
ENTRYPOINT ["/admin"]
