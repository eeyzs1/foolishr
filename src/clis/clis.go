package clis

import (
	"github.com/eeyzs1/foolishr/src/config"
	"github.com/eeyzs1/foolishr/src/router"

	"fmt"
)

func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Clean() {

	fmt.Println("Clean~~~~~~~~~~~~~")

}



