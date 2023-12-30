# syntax=docker/dockerfile:1
FROM golang:1.21 as build
WORKDIR /src

COPY go.mod /src
COPY go.sum /src
COPY main.go /src

RUN go build -o /bin/access-www ./main.go

FROM ubuntu:22.04

RUN mkdir -p /usr/share/access-www/assets && mkdir -p /usr/share/access-www/templates && mkdir -p /usr/share/access-www/ssl

RUN apt-get update && apt install openssl
RUN openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -subj \
    "/C=US/CN=localhost" \
    -keyout /tmp/access.baf.key -out /tmp/access.baf.crt && \
    ls -la /tmp && \
    cp /tmp/access.baf.key /usr/share/access-www/ssl/ && \
    cp /tmp/access.baf.crt /usr/share/access-www/ssl/

COPY assets /usr/share/access-www/assets/
COPY templates /usr/share/access-www/templates/
COPY --from=build /bin/access-www /bin/access-www

RUN ls /bin
CMD ["/bin/access-www"]