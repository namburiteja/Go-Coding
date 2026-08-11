# Build from repo root:
# docker build -f services/report/Dockerfile -t report-service:v1 .

FROM golang:1.26-alpine AS build

WORKDIR /repo

COPY shared ./shared
COPY services/report ./services/report

WORKDIR /repo/services/report

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/report ./cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /bin/report .

EXPOSE 9095

ENTRYPOINT ["./report"]