# Build from repo root:
# docker build -f services/merchant/Dockerfile -t merchant-service:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY services/merchant ./services/merchant

WORKDIR /repo/services/merchant

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/merchant ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/merchant .

EXPOSE 9092

ENTRYPOINT ["./merchant"]