package db_management

import (
	"fmt"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"net"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func init_mysql(){

}

func init_redis(){

}


func op() {
	dsn := "/path/to/mysql.sock"  // 请替换为实际的 MySQL 套接字路径
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "unix(" + dsn + ")/dbname?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
		return
	}
	defer db.Close()


	conn, err := net.Dial("unix", "/tmp/example.sock")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected to /tmp/example.sock")


	client := redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:    "/path/to/your/redis.sock",
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)





	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 迁移数据库表
	db.AutoMigrate(&User{})

	// 操作关系型数据库
	user := User{Name: "John", Email: "john@example.com"}
	db.Create(&user)
	fmt.Println("User created:", user)

	// 连接到 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// 操作 Redis
	err = rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic("Failed to set key in Redis")
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic("Failed to get key from Redis")
	}
	fmt.Println("Value from Redis:", val)
}