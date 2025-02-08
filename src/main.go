package main

import (
	"fmt"
	"os"
	"os/exec"
	// "flag"
	
	"foolishr/src/config"
	"foolishr/src/router"
	// "path/filepath"

	// "github.com/eeyzs1/foolishr/src/clis"
)

func main() {
	// 检查参数数量 
    if len(os.Args) < 2 {
        fmt.Println("错误：未提供子命令")
        return 
    }
 
    command := os.Args[1]
    switch command {
    case "start":
        fmt.Println("启动服务...")
        // 执行启动逻辑 
    case "stop":
        fmt.Println("停止服务...")
        // 执行停止逻辑 
    default:
        fmt.Printf("未知命令: %s\n", command)
    }

	// go build -o foolishr ./src/main.go
	// ./foolishr
	// 定义标志参数（可选）foolishr start -port=9000  # 输出 "启动服务，端口: 9000"
	// foolishr stop              # 输出 "停止服务"
	// didnt use flag coz too many args its better to use config file
	//  port := flag.Int("port", 8080, "服务端口号")//Int defines an int flag with specified name, default value, and usage string.
	//  flag.Parse()

	// var (
	// 	name    = flag.String("name", "guest", "用户名称")
	// 	age     = flag.Int("age", 0, "用户年龄")
	// 	isVIP   = flag.Bool("vip", false, "是否VIP用户")
	// 	timeout = flag.Duration("timeout", 5*time.Second, "请求超时时间")
	// )
	
	// flag.Parse() // 必须调用解析方法 
	
	// // 输出解析结果 
	// fmt.Printf("用户信息:\n 姓名:%s\n 年龄:%d\n VIP:%t\n 超时:%v\n",
	// 	*name, *age, *isVIP, *timeout)
	
	// // 获取非标志参数 
	// fmt.Println("\n其他参数:", flag.Args())
	
	//  // 获取位置参数 
	//  args := flag.Args()
	//  if len(args) < 1 {
	// 	 fmt.Println("错误：未提供子命令")
	// 	 return 
	//  }
  
	//  switch args[0] {
	//  case "start":
	// 	 fmt.Printf("启动服务，端口: %d\n", *port)
	//  case "stop":
	// 	 fmt.Println("停止服务")
	//  default:
	// 	 fmt.Printf("未知命令: %s\n", args[0])
	//  }
  
	// 	cookie, err := c.Cookie("gin_cookie")
  
	// 	if err != nil {
	// 		cookie = "NotSet"
	// 		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	// 	}
  
	// 	fmt.Printf("Cookie value: %s \n", cookie)
	// })
  
	// router.Run()

	// cmd := exec.Command("killall", "-TERM", "fooolishr")
    // err := cmd.Run()
    // if err != nil {
    //   fmt.Println("Error executing restart command:", err)
    //   return
    // }

  }


func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Stop() {
	cmd := exec.Command("killall", "-TERM", "fooolishr")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing restart command:", err)
		return
	}
}
