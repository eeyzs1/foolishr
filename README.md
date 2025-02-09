# foolishr

This project aims to record my study of web dev based on go, grpc,gin,gorm,vue,docker,k8s


## tool used:
- vscode, extensions:vscode-proto3, clang-format,go,docker
- chocolatey


##prepare grpc:
install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

##learning material:
- https://dev.to/techschoolguru
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