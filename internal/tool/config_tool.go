package tool

import (
	"fmt"
	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()
	// 文件所在目录
	config.AddConfigPath("./tool/")
	// 文件名
	config.SetConfigName("tool")
	// 文件类型
	config.SetConfigType("toml")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件")
		} else {
			fmt.Println("配置文件出错")
		}
	}
}
func GetConfig() *viper.Viper {
	return config
}
