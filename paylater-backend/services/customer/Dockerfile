# build from repo root: docker build -f services/customer/Dockerfile .
FROM golang:1.26-alpine AS build
WORKDIR /repo
COPY shared ./shared
COPY services/customer ./services/customer
WORKDIR /repo/services/customer
RUN go mod download && CGO_ENABLED=0 go build -o /bin/customer ./cmd
FROM alpine:3.20
COPY --from=build /bin/customer /customer
EXPOSE 9093
ENTRYPOINT ["/customer"]
