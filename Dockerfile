FROM alpine:latest

COPY / /
WORKDIR /

CMD ["./app"]