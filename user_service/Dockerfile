FROM golang:1.24-alpine
# ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,https://gocenter.io,direct

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN apk add make
RUN make build

CMD ["./build/app"]