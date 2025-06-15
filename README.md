# foolishr

This project aims to record my study of web dev based on go, grpc,gin,gorm


## tools used:
- vscode, extensions:vscode-proto3, go
- chocolatey
- postman, swagger.io

## how to run:
- cd src
- make gen to generate the protocol buff related go files and swagger realted json files
- for tls, run gen.sh under cert folder
- run grpc part: make server(-tls), make client(-tls)
- for load balance: nginx, and check the nginx.conf for ssl, grpc services mapping; then run multiple server by make server(1/2)(-tls)
- for restful services with grpc gateway: make rest


##prepare grpc:
install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

##learning material:
- https://dev.to/techschoolguru
- https://github.com/techschool
- https://grpc.org.cn/docs/languages/go/basics/
- https://protobuf.dev/
- https://gorm.io/zh_CN/docs/models.html
- https://grpc.io/
- https://protobuf.dev/reference/go/go-generated/#package

package mypkg 

var PublicVar int = 42  // 可被外部包访问 
var privateVar int = 10  // 仅限本包使用 

func PublicFunc() {}     // 可被外部调用 
func privateFunc() {}    // 仅限本包调用 
main包是程序入口，但无法被其他包导入

test pkg name should be name should be the pkg+"_test"
after run pkg test, can find blue label source code is covered code, red is uncovered
go test pkgname, example: go test ./serializer
go test ./... run unit test on all sub pkgs
