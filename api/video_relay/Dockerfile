FROM golang:1.18-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o /app .

# Prod container
FROM gcr.io/distroless/static as prod

COPY --from=builder /app /bin/app

USER 10001:10001

CMD ["/bin/app"]