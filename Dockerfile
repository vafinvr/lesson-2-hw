FROM golang:latest AS builder
ADD . /src
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app .

FROM alpine:latest
COPY --from=builder /app /
RUN chmod +x /app
EXPOSE 8000
CMD ["/app"]
