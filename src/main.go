package main
import (
	"fmt"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"net"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/go-redis/redis"
	"github.com/gin-gonic/gin"
  )
  
func main() {
	router := gin.Default()
  
	router.GET("/cookie", func(c *gin.Context) {
  
		cookie, err := c.Cookie("gin_cookie")
  
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
  
		fmt.Printf("Cookie value: %s \n", cookie)
	})
  
	router.Run()

	cmd := exec.Command("killall", "-HUP", "appweiyigeek")
    err := cmd.Run()
    if err != nil {
      fmt.Println("Error executing restart command:", err)
      return
    }

  }