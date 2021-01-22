
DOCKER_TAG="dormao/fireplace:1.1.1"

# create bin dir
.PHONY:dir
dir:
	mkdir -p ./bin

# build amd64
.PHONY:amd64
amd64: 
	GOOS=linux GOARCH=amd64 go build -o bin/fireplace-server main.go

# build arm64
.PHONY:arm64
arm64: 
	GOOS=linux GOARCH=arm64 go build -o bin/fireplace-server main.go

# build amd64 docker
.PHONY:amd64-docker
amd64-docker: amd64
	docker build -t ${DOCKER_TAG} -f ./Dockerfile .

# build arm64 docker
.PHONY:arm64-docker
arm64-docker: arm64
	docker build -t ${DOCKER_TAG} -f ./Dockerfile .