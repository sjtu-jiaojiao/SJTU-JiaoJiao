mkdir build
cp config.json ./build

docker build . -t sjtujj
docker build ./api/auth -t sjtujj/api-auth
docker build ./api/user -t sjtujj/api-user
docker build ./api/sellinfo -t sjtujj/api-sellinfo
# docker api insert before this
docker build ./srv/auth -t sjtujj/srv-auth
docker build ./srv/user -t sjtujj/srv-user
docker build ./srv/sellinfo -t sjtujj/srv-sellinfo
# docker srv insert before this

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-auth ./api/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-user ./api/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-sellinfo ./api/sellinfo/main.go
# build api insert before this
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-auth ./srv/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-user ./srv/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-sellinfo ./srv/sellinfo/main.go
# build srv insert before this