FROM golang:alpine AS builder
RUN apk add --no-cache make git
WORKDIR /app/
COPY auth_service/go.mod auth_service/go.sum ./
COPY auth_service/ ./
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app
COPY auth_service/config/config.yaml /config/config.yaml
WORKDIR /
CMD ["/app"]