FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app ./
RUN chmod +x ./app
EXPOSE 8080
ENTRYPOINT ["./app"]

