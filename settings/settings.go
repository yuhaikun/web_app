package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config") // 指定配置文件名称（不需要后缀）
	//viper.SetConfigType("yaml")   // 指定配置文件类型（专指用于从远程获取配置信息时指定配置）
	viper.AddConfigPath(".") // 指定查找配置文件的路径(相对路径)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInconfig() failed,err:%v\n", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//r := gin.Default()
	//if err := r.Run(fmt.Sprintf(":%d", viper.Get("port"))); err != nil {
	//	panic(err)
	//}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
