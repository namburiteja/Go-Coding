# Build from repo root:
# docker build -f services/ledger/Dockerfile -t ledger-service:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY services/ledger ./services/ledger

WORKDIR /repo/services/ledger

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/ledger ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/ledger .

EXPOSE 9094

ENTRYPOINT ["./ledger"]