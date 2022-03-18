FROM golang:1.17-buster as builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o server ./cmd

FROM debian:buster
COPY --from=builder /app/server /app/server
COPY ./wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh
EXPOSE 9000
ENTRYPOINT ["/app/wait-for-it.sh", "http://editor-db-alpha:9080", "--", "/app/server"]
