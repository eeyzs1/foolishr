package main

import (
	// "fmt"
	"github.com/eeyzs1/foolishr/src/clis"
  )
  
func main() {

	defer clis.Clean()
	clis.Start()

	// router := gin.Default()
  
	// router.GET("/cookie", func(c *gin.Context) {
  
	// 	cookie, err := c.Cookie("gin_cookie")
  
	// 	if err != nil {
	// 		cookie = "NotSet"
	// 		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	// 	}
  
	// 	fmt.Printf("Cookie value: %s \n", cookie)
	// })
  
	// router.Run()

	// cmd := exec.Command("killall", "-HUP", "appweiyigeek")
    // err := cmd.Run()
    // if err != nil {
    //   fmt.Println("Error executing restart command:", err)
    //   return
    // }

  }