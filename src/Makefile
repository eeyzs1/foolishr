gen:
	protoc -I proto proto/*.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=:swagger

# 定义操作系统敏感的操作 
ifeq ($(OS),Windows_NT)
  RM := del /Q /F 
  MKDIR := mkdir 
else 
  RM := rm -f 
  MKDIR := mkdir -p 
endif 
 
clean:
	@echo "[INFO] 清理生成文件..."
	$(RM) pb\*.go 
	$(RM) swagger\*

run:
	go run main.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server:
	go run cmd/server/main.go -port 8080

server-tls:
	go run cmd/server/main.go -port 8080 -tls

rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 127.0.0.1:8080

client:
	go run cmd/client/main.go -address 127.0.0.1:8080

client-tls:
	go run cmd/client/main.go -address 127.0.0.1:8080 -tls

test:
	go test -cover -race ./... #race to check race condition, cover measure the code coverage of test

cert:
	cd cert; ./gen.sh; cd ..

# 在 Makefile 中，.PHONY: clean gen server client test cert 用于声明一系列伪目标（Phony Targets），其核心作用是通过显式标记避免目标名与文件名冲突，并优化构建逻辑
# 防止目标名与文件名冲突
# 当目录中存在与目标同名的文件（如 clean、test）时，Make 默认会认为该目标对应的文件已存在，导致命令不执行。通过 .PHONY 标记可强制声明这些目标是操作指令而非文件，确保每次执行 make <target> 时命令都会运行
# # 若未声明 .PHONY 且存在 `clean` 文件，则 make clean 会失效 
# .PHONY: clean 
# clean:
#     rm -rf *.o 
# 提高执行效率
# Make 默认会检查目标依赖文件的更新时间，但对伪目标（如 server、client）会跳过依赖检查，直接执行命令，减少不必要的计算

.PHONY: clean gen server client test cert 