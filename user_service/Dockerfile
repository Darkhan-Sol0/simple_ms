FROM golang:alpine AS builder
RUN apk add --no-cache make git
WORKDIR /app/user_service
COPY user_service/go.mod user_service/go.sum ./
COPY user_service/ ./
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/user_service/app /app
COPY user_service/config/config.yaml /config/config.yaml
WORKDIR /
CMD ["/app"]