FROM golang:1.23-alpine AS builder
RUN mkdir -p /app
WORKDIR /app
COPY server .
RUN go build -o app .


FROM alpine:3
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
ENTRYPOINT [ "./app" ]
