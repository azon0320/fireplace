
DOCKER_TAG="dormao/fireplace:1.1.1"

# create bin dir
.PHONY:dir
dir:
	mkdir -p ./bin

# build amd64
.PHONY:amd64
amd64: 
	GOOS=linux GOARCH=amd64 go build -o bin/fireplace-server main.go

# build arm
.PHONY:arm
arm: 
	GOOS=linux GOARCH=arm go build -o bin/fireplace-server main.go

# build amd64 docker
.PHONY:amd64-docker
amd64-docker: amd64
	docker build -t ${DOCKER_TAG} -f ./Dockerfile . && rm -rf ./bin/*

# build arm docker
.PHONY:arm-docker
arm-docker: arm
	docker build -t ${DOCKER_TAG} -f ./Dockerfile . && rm -rf ./bin/*