read -p 'Service name: ' srvname
srvnamel=${srvname,,}
cp -R ./template/api ./api/$srvnamel
cp -R ./template/srv ./srv/$srvnamel

sed -i "s/{{SERVICE_LNAME}}/$srvnamel/g" ./api/$srvnamel/Dockerfile
sed -i "s/{{SERVICE_NAME}}/$srvname/g" ./api/$srvnamel/main.go
mv ./srv/$srvnamel/proto/NAME.proto ./srv/$srvnamel/proto/$srvnamel.proto 
sed -i "s/{{SERVICE_LNAME}}/$srvnamel/g" ./srv/$srvnamel/Dockerfile
sed -i "s/{{SERVICE_NAME}}/$srvname/g" ./srv/$srvnamel/main.go

sed -i "/# docker api insert before this/i sudo docker build ./api/$srvnamel -t sjtujj/api-$srvnamel" build.sh
sed -i "/# docker srv insert before this/i sudo docker build ./srv/$srvnamel -t sjtujj/srv-$srvnamel" build.sh
sed -i "/# build api insert before this/i CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/api-$srvnamel ./api/$srvnamel/main.go" build.sh
sed -i "/# build srv insert before this/i CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/srv-$srvnamel ./srv/$srvnamel/main.go" build.sh
sed -i "/# docker api insert before this/i sudo docker push sjtujj/api-$srvnamel" deploy.sh
sed -i "/# docker srv insert before this/i sudo docker push sjtujj/srv-$srvnamel" deploy.sh

echo "- name: api-$srvnamel" >> .realize.yaml
echo "  path: api/$srvnamel" >> .realize.yaml
cat<<EOF >> .realize.yaml
  args:
    - --registry=consul
  commands:
    run:
      status: true
  watcher:
    extensions:
      - go
    paths:
      - /
      - ../../utils
    scripts:
    - command: apidoc
      type: after
      path: ../../
      output: true
EOF

echo "- name: srv-$srvnamel" >> .realize.yaml
echo "  path: srv/$srvnamel" >> .realize.yaml
cat<<EOF >> .realize.yaml
  args:
  - --registry=consul
  commands:
    run:
      status: true
  watcher:
    extensions:
    - proto
    - go
    paths:
    - /
    - ../../utils
    scripts:
    - command: bash generate_proto.sh
      type: before
      path: ../../
      output: true
    - command: apidoc
      type: after
      path: ../../
      output: true
EOF