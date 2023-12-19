# syntax=docker/dockerfile:1
FROM golang:1.21 as build
WORKDIR /src

COPY go.mod /src
COPY go.sum /src
COPY main.go /src
RUN go build -o /bin/hello ./main.go

FROM ubuntu:22.04
COPY --from=build /bin/hello /bin/hello
RUN ls /bin
CMD ["/bin/hello"]