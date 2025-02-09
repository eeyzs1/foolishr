gen:
	protoc proto/*.proto --go_out=.

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

run:
	go run main.go


