package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


type FnRegistRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []FnRegistRouter
)

func RegisterRouter(fn FnRegistRouter){
	if fn == nil { return }
	gfnRouters = append(gfnRouters, fn)
}

func InitRouter() {
	r := gin.Default()

	rgPublic := r.Group("/api/v/1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRouters()

	for _, FnRegisterRouter := range gfnRouters {
		FnRegisterRouter(rgPublic, rgAuth)
	}

	stPort := viper.GetString("server.port")

	if stPort == "" {
		stPort = "10086"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))

	if err != nil {
		panic(fmt.Sprintf("Start Server Error: %s", err.Error()))
	}

}


func InitBasePlatformRouters() {
	InitUserRouters()

}

func InitUserRouters() {
	RegisterRouter(func(
		rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})

		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id":1, "name": "whoever"},
					{"ps": "whatever"},
				},
			})
		})

		rgAuthUser.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id": 1,
				"name": "whoever",
			})
		})

	})
}





