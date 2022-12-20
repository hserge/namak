## Build
FROM golang:1.19-alpine AS builder
WORKDIR /build
# Fetch dependencies and build
COPY ../go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /nmk-app
# ENTRYPOINT ["tail", "-f", "/dev/null"]

## Deploy
# FROM scratch
FROM alpine:latest
# set permission to "nobody"
COPY --chown=65534:65534 --from=builder /namak-app /namak
USER 65534
EXPOSE 8080
ENTRYPOINT ["/namak"]
# ENTRYPOINT ["tail", "-f", "/dev/null"]