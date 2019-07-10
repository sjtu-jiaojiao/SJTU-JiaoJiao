FROM golang:latest AS build
WORKDIR /
ADD . .

ENV GOFLAGS -mod=vendor

RUN ./build.sh