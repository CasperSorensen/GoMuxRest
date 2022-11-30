FROM golang:alpine

Run mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 9090

ENTRYPOINT ["/var/www/app"]