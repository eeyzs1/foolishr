package router

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
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

	srv := &http.Server{
		Addr:    stPort,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	// stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
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





