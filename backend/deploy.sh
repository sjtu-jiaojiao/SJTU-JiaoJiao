echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

docker push sjtujj/api-auth
docker push sjtujj/api-user
docker push sjtujj/api-sellinfo
# docker api insert before this
docker push sjtujj/srv-auth
docker push sjtujj/srv-user
docker push sjtujj/srv-sellinfo
# docker srv insert before this

ssh -i private.key -p 30710 centos@202.120.40.8 bash update.sh