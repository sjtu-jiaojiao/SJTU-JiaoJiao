echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

sudo docker push sjtujj/api-auth
sudo docker push sjtujj/api-user
sudo docker push sjtujj/api-sellinfo
# docker api insert before this
sudo docker push sjtujj/srv-auth
sudo docker push sjtujj/srv-user
sudo docker push sjtujj/srv-sellinfo
# docker srv insert before this

ssh -i private.key -p 30710 centos@202.120.40.8 bash update.sh