docker build . -t sjtujj
docker build ./api/auth -t sjtujj/api-auth
docker build ./api/user -t sjtujj/api-user
docker build ./api/sellinfo -t sjtujj/api-sellinfo
sudo docker build ./api/buyinfo -t sjtujj/api-buyinfo
sudo docker build ./api/file -t sjtujj/api-file
# docker api insert before this
docker build ./srv/auth -t sjtujj/srv-auth
docker build ./srv/user -t sjtujj/srv-user
docker build ./srv/sellinfo -t sjtujj/srv-sellinfo
sudo docker build ./srv/buyinfo -t sjtujj/srv-buyinfo
sudo docker build ./srv/file -t sjtujj/srv-file
# docker srv insert before this 
