package common

import "github.com/spf13/viper"

type ServerProperties struct {
	ErShouFang string
	ChengJiao  string
}

var Properties = &ServerProperties{}

func InitConfig() {
	Properties.ErShouFang = viper.GetString("url.ershoufang")
	Properties.ChengJiao = viper.GetString("url.chengjiao")
}
