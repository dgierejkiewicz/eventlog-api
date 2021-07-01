# Dockerfile args
ARG APP_VERSION=0.0.1

FROM golang:1.13-alpine3.11 AS builder
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd

FROM alpine:latest AS production
COPY --from=builder /bin/app /bin/app
EXPOSE "3000"
ENTRYPOINT [ "/bin/app" ]