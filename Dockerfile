FROM golang:1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy -compat 1.17
RUN go build -o app.bin cmd/api/main.go

FROM alpine:3 AS runner
WORKDIR /app
COPY --from=builder /app/app.bin app.bin
RUN chown -R nobody: /app
ENV HOST=0.0.0.0
ENV PORT=8080
EXPOSE 8080
CMD ["./app.bin"]
