package config

import (
	"fmt"
	// "os"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	// "gopkg.in/yaml.v3"
)

// type Config struct {
//     App struct {
//         Name string `mapstructure:"name"`
//     } `mapstructure:"app"`
//     Server struct {
//         Port int `mapstructure:"port"`
//     } `mapstructure:"server"`
//     Debug bool `mapstructure:"debug"`
// }

var (
	ViperInst = viper.New()//actually not needed if only has 1 viper instance that viper pkg provided 1 global viper inst, leave here just for possible extention example
	// Cfg Config 
)

func InitConfig() {
	ViperInst.SetConfigName("configs")
	ViperInst.SetConfigType("yml")
	ViperInst.AddConfigPath("./")
	
	if err := ViperInst.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}

    // if err := ViperInst.Unmarshal(&Cfg); err != nil {
    //     fmt.Printf("配置解析失败: %v", err)
    // }

	ViperInst.Set("database.host",  "127.0.0.1")    // 单级键 
    ViperInst.Set("redis.port",  6379)            // 多级嵌套键[4]()[11]()
 
    // 保存配置到文件 
    if err := ViperInst.WriteConfig(); err != nil {  // 使用预定义路径[2]()
        panic(fmt.Errorf("配置保存失败: %w", err))
    }
    
    fmt.Println("配置更新成功")

	fmt.Println(ViperInst.GetString("server.port"))

	ViperInst.WatchConfig()
	ViperInst.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更新:", e.Name)
		//todo: 重新加载配置逻辑...
	})

}

func GetServerPort() string {
	return ViperInst.GetString("server.port")
}

// func SaveConfig(){
// 	allSettings := ViperInst.AllSettings()
// 	if data, err := yaml.Marshal(allSettings);err!=nil{
// 		fmt.Printf("配置存储失败: %v", err)
// 	}
// 	 // 写入文件（权限0644：用户可读写，其他只读）
// 	 err = os.WriteFile("config.yaml",  data, 0644)
// 	 if err != nil {
// 		 panic("写入文件失败: " + err.Error())
// 	 }
// }
