docker build . -t sjtujj
docker build ./api/auth -t sjtujj/api-auth
docker build ./api/user -t sjtujj/api-user
docker build ./api/sellinfo -t sjtujj/api-sellinfo
# docker api insert before this
docker build ./srv/auth -t sjtujj/srv-auth
docker build ./srv/user -t sjtujj/srv-user
docker build ./srv/sellinfo -t sjtujj/srv-sellinfo
# docker srv insert before this 
