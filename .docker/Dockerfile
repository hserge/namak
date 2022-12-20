## Build
FROM golang:1.19-alpine AS build
WORKDIR /app
COPY .. .
RUN cd ./src
RUN ls -la
RUN pwd
RUN go mod download
ENTRYPOINT ["tail", "-f", "/dev/null"]