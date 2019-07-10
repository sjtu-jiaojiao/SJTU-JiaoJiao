FROM golang:latest AS build
WORKDIR /
ADD . .

ENV GOFLAGS -mod=vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-auth ./api/auth/main.go && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-user ./api/user/main.go && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-auth ./srv/auth/main.go && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-user ./srv/user/main.go