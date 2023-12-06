package main

import (
	// "fmt"
	"fmt"
	"os"
	"os/exec"
	// "path/filepath"

	// "github.com/eeyzs1/foolishr/src/clis"
)
  
func main() {

	currentDirectory, err := os.Getwd() //todo:just for go run quick test, later change to go build, run server with clis then remove this line
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current directory:", currentDirectory)
	targetFile := currentDirectory + "/clis/clis.go"

	cmd := exec.Command("go", "run","-exec=foolishr", targetFile, "Start")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing start command:", err)
		fmt.Println("Command output:", string(output))
		return
	}

	// go build ./main.go -o foolishr
	// ./foolishr
  
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