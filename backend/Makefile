SHELL = /bin/bash
OUTPUTDIR = ./build
SRVLIST = ./srvlist.list

.ONESHELL:
.PHONY: build docker deps clean deploy proto add doc test help

all: help

## build: Build binary file.
build:
	@export GOFLAGS=-mod=vendor
	@while read -r line;do 				\
		echo Building api-$$line...;	\
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUTPUTDIR)/api-$$line ./api/$$line/main.go;	\
		echo Building srv-$$line...;	\
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUTPUTDIR)/srv-$$line ./srv/$$line/main.go;	\
	done < "$(SRVLIST)"
	@echo Building api-tag...
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUTPUTDIR)/api-tag ./api/tag/main.go
	@cp config.json $(OUTPUTDIR)
	@echo Success

## proto: Generate proto file.
proto:
	@echo Generating proto file...
ifdef SRV
	@for file in `find ./srv/$(SRV) -name "*.proto"`;do	 \
		dir=$$(dirname "$${file}");					\
		protoc --proto_path=$$dir --micro_out=$$dir --go_out=$$dir `basename "$${file}"`;	\
	done
else
	@for file in `find ./srv/ -name "*.proto"`;do	 \
		dir=$$(dirname "$${file}");					\
		protoc --proto_path=$$dir --micro_out=$$dir --go_out=$$dir `basename "$${file}"`;	\
	done
endif
	@echo Success

## docker: Build docker images.
docker: build
	@sudo docker build . -t sjtujj
	@sudo docker build ./api/tag -t sjtujj/api-tag
	@while read -r line;do		\
		sudo docker build ./api/$$line -t sjtujj/api-$$line;	\
		sudo docker build ./srv/$$line -t sjtujj/srv-$$line;	\
	done < "$(SRVLIST)"

## deploy: Deploy to server.
deploy:
	@echo "$$DOCKER_PASSWORD" | docker login -u "$$DOCKER_USERNAME" --password-stdin
	@sudo docker push sjtujj/api-tag
	@while read -r line;do		\
		sudo docker push sjtujj/api-$$line;	\
		sudo docker push sjtujj/srv-$$line;	\
	done < "$(SRVLIST)"
	@openssl aes-256-cbc -K $$encrypted_3b095c6852fb_key -iv $$encrypted_3b095c6852fb_iv -in private.key.enc -out private.key -d
	@chmod 400 private.key
	@ssh -i private.key -p 30710 centos@202.120.40.8 bash service.sh update
	@echo Success

## add: Add new service.
add:
	@read -p 'Service name: ' srvname
	@srvnamel=`echo "$$srvname" | tr '[:upper:]' '[:lower:]'`

	@cp -R ./template/api ./api/$$srvnamel
	@cp -R ./template/srv ./srv/$$srvnamel

	@mv ./api/$$srvnamel/main_test.tpl ./api/$$srvnamel/main_test.go
	@mv ./api/$$srvnamel/main.tpl ./api/$$srvnamel/main.go
	@mv ./srv/$$srvnamel/main_test.tpl ./srv/$$srvnamel/main_test.go
	@mv ./srv/$$srvnamel/main.tpl ./srv/$$srvnamel/main.go
	@mv ./srv/$$srvnamel/mock/mock.tpl ./srv/$$srvnamel/mock/mock.go
	@mv ./srv/$$srvnamel/proto/NAME.proto ./srv/$$srvnamel/proto/$$srvnamel.proto

	@sed -i "s/{{SERVICE_LNAME}}/$$srvnamel/g" ./api/$$srvnamel/Dockerfile
	@sed -i "s/{{SERVICE_NAME}}/$$srvname/g" ./api/$$srvnamel/main.go
	@sed -i "s/{{SERVICE_LNAME}}/$$srvnamel/g" ./srv/$$srvnamel/Dockerfile
	@sed -i "s/{{SERVICE_NAME}}/$$srvname/g" ./srv/$$srvnamel/main.go

	@echo $$srvnamel >> $(SRVLIST)

	@cat ./template/.realize.yaml >> .realize.yaml
	@sed -i "s/{{SERVICE_LNAME}}/$$srvnamel/g" .realize.yaml

## test: Run go test.
test:
	@go list -f '{{if len .TestGoFiles}}"s=`echo {{.Dir}} | tail -c 6` && [ $$s == "utils" ] && go test -tags=test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}} || ([ $$s != "stest" ] && go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}} || :)"{{end}}' ./... | xargs -L 1 bash -c

## doc: Generate document.
doc:
	@apidoc -i api -i srv

## clean: Clean build files.
clean:
	-@go clean
	-@rm -rf $(OUTPUTDIR)
	-@find . -name .coverprofile | xargs -r rm

## deps: Install dependences.
deps:
	go mod download
	@export GO111MODULE=off
	go get github.com/micro/micro
	go get github.com/smartystreets/goconvey
	go get github.com/oxequa/realize
	go get github.com/micro/protoc-gen-micro
	go get github.com/golang/protobuf/protoc-gen-go
	npm install apidoc -g

## help: Show this help.
help: Makefile
	@echo Usage: make [command]
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'