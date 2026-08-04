# build from repo root: docker build -f services/merchant/Dockerfile .
FROM golang:1.26-alpine AS build
WORKDIR /repo
COPY shared ./shared
COPY services/merchant ./services/merchant
WORKDIR /repo/services/merchant
RUN go mod download && CGO_ENABLED=0 go build -o /bin/merchant ./cmd
FROM alpine:3.20
COPY --from=build /bin/merchant /merchant
EXPOSE 9092
ENTRYPOINT ["/merchant"]
