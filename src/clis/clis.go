package clis

import (
	"os/exec"

	"github.com/eeyzs1/foolishr/src/config"
	"github.com/eeyzs1/foolishr/src/router"

	"fmt"
)

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



