package router

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"foolishr/src/config"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type FnRegistRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

type User struct{
	gorm.Model
	Name string "gorm:"
}

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-"
	nameLength = 10
	letterIdxBits = 6                      // 64字符需6位(2^6=64)
    letterIdxMask = 1<<letterIdxBits - 1   // 掩码(0b111111)
	moveLength = 5							//5*10 = 50 + 6<64
)

var (
	gfnRouters []FnRegistRouter
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	rMu sync.Mutex
)

func RegisterRouter(fn FnRegistRouter) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

func InitRouter() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	r := gin.Default()

	rgPublic := r.Group("/api/public")
	rgAuth := r.Group("/api/auth")

	InitBasePlatformRouters()

	for _, FnRegisterRouter := range gfnRouters {
		FnRegisterRouter(rgPublic, rgAuth)
	}

	stPort := config.GetServerPort()

	if stPort == "" {
		stPort = "10086"
	}

	srv := &http.Server{
		Addr:    ":" + stPort,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// block here to waiting for shutdown signal
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

	log.Println("Server exited")
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
		rgAuthUser.GET("register", func(ctx *gin.Context) {
			// get parameters
			name := ctx.PostForm("name")
			telephone := ctx.PostForm("telephone")
			password := ctx.PostForm("password")

			// show the input check
			//todo:put calculation in the front end
			if len(telephone) != 11{
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"length of phone num should be 11"})
			}
			if len(name) == 0{
				name = RadnomString()
			}

			
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "whoever"},
					{"ps": "whatever"},
				},
			})
		})

		rgAuthUser.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "whoever",
			})
		})

	})
}


func RadnomString() string{
	result := make([]byte,nameLength)
	rMu.Lock()
    defer rMu.Unlock()
	randomNum := r.Int63()
	var Mask int64 = letterIdxMask
	for i := range nameLength {
		result[i] = letters[randomNum&(Mask)]
		Mask <<= moveLength
	}
	return string(result)
}
