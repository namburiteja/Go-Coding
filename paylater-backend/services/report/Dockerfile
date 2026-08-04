# build from repo root: docker build -f services/report/Dockerfile .
FROM golang:1.26-alpine AS build
WORKDIR /repo
COPY shared ./shared
COPY services/report ./services/report
WORKDIR /repo/services/report
RUN go mod download && CGO_ENABLED=0 go build -o /bin/report ./cmd
FROM alpine:3.20
COPY --from=build /bin/report /report
EXPOSE 9095
ENTRYPOINT ["/report"]
