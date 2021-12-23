# Kratos Project Template
学习 kratos 项目时写的练手项目

- main 分支，路由模块使用的是 Kratos 提供的 http 模块，为主分支
- kratos-http-template 分支，路由模块使用的是 Kratos 提供的 http 模块
- kratos-gin-template  分支，路由模块使用的是 Kratos 提供的 http 模块


## Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API swagger json files by proto file
make swagger
# Generate API validator files by proto file
make validate
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

