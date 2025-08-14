FROM golang:alpine AS builder
RUN apk add --no-cache make git
WORKDIR /app/
COPY monolit_service/go.mod monolit_service/go.sum ./
COPY monolit_service/ ./
RUN make build

FROM alpine:latest
RUN apk --no-cache add make go ca-certificates
COPY --from=builder /app/app /app
COPY monolit_service/config/config.yaml /config/config.yaml
COPY monolit_service/migrations/ /migrations/
COPY monolit_service/Makefile/ /Makefile
WORKDIR /
CMD ["sh", "-c", "make init && make goose_install && make goose_up && /app"]