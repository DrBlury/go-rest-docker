FROM golang:latest AS builder

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
######################################
FROM alpine:latest
WORKDIR /root/
LABEL AUTHOR="JULIAN BENSCH"
COPY /public ./public
COPY defaults.yml .
COPY --from=builder /app/app .
CMD ["./app"]