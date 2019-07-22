docker build . -t sjtujj
docker build ./api/auth -t sjtujj/api-auth
docker build ./api/user -t sjtujj/api-user
docker build ./api/sellinfo -t sjtujj/api-sellinfo
sudo docker build ./api/file -t sjtujj/api-file
sudo docker build ./api/avatar -t sjtujj/api-avatar
sudo docker build ./api/content -t sjtujj/api-content
# docker api insert before this
docker build ./srv/auth -t sjtujj/srv-auth
docker build ./srv/user -t sjtujj/srv-user
docker build ./srv/sellinfo -t sjtujj/srv-sellinfo
sudo docker build ./srv/file -t sjtujj/srv-file
sudo docker build ./srv/avatar -t sjtujj/srv-avatar
sudo docker build ./srv/content -t sjtujj/srv-content
# docker srv insert before this 
