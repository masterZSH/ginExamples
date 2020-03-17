package config

import "github.com/spf13/viper"
import "fmt"

type Config struct {
	Env   string `json:"env"`
	Mysql Mysql  `json:"mysql"`
}

type Mysql struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

var SysConfig Config

func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")     // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&SysConfig); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", SysConfig)
}
