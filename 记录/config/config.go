package config

import (
	"github.com/spf13/viper"
)

type MyConfig struct {
	Ip      string
	Port    string
	Bottoms []int
}

var Config = MyConfig{}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./记录/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败,err = " + err.Error())
	}
	Config.Bottoms = viper.GetIntSlice("Game.Bottoms")
	Config.Ip = viper.GetString("Server.Ip")
	Config.Port = viper.GetString("Server.Port")
}
