package clis

import (
	"config"
	"fmt"
	"router"
)

func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Clean() {

	fmt.Println("Clean~~~~~~~~~~~~~")

}



