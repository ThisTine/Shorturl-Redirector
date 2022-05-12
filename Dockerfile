FROM golang:alpine AS builder

WORKDIR /app

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine

WORKDIR /app


COPY --from=builder /app/shorturl-redirector .


EXPOSE 8000

ENTRYPOINT [ "/app/shorturl-redirector" ]