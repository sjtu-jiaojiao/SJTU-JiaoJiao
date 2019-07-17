mkdir build
cp config.json ./build

sudo docker build . -t sjtujj
sudo docker build ./api/auth -t sjtujj/api-auth
sudo docker build ./api/user -t sjtujj/api-user
sudo docker build ./api/sellinfo -t sjtujj/api-sellinfo
# docker api insert before this
sudo docker build ./srv/auth -t sjtujj/srv-auth
sudo docker build ./srv/user -t sjtujj/srv-user
sudo docker build ./srv/sellinfo -t sjtujj/srv-sellinfo
# docker srv insert before this

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-auth ./api/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-user ./api/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-sellinfo ./api/sellinfo/main.go
# build api insert before this
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-auth ./srv/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-user ./srv/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-sellinfo ./srv/sellinfo/main.go
# build srv insert before this