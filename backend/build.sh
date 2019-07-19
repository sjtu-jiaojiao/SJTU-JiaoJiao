CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-auth ./api/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-user ./api/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-sellinfo ./api/sellinfo/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-buyinfo ./api/buyinfo/main.go
# build api insert before this
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-auth ./srv/auth/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-user ./srv/user/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-sellinfo ./srv/sellinfo/main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-buyinfo ./srv/buyinfo/main.go
# build srv insert before this

cp config.json ./build