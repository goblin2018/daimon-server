package conf

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	App   App   `json:"app"`
	DB    DB    `json:"db"`
	Redis Redis `json:"redis"`
}

type App struct {
	Mode      string `json:"mode"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Key       string `json:"key"`
	LogStd    bool
	Log       string `json:"log"`
	BlackList []string
}

type DB struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var C Config

func init() {
	pflag.String("env", "dev", "server runing env ")
	pflag.Bool("logstd", true, "log to console")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	env := viper.GetString("env")

	viper.AddConfigPath(getConfPath())
	viper.SetConfigName("app." + env)
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config error: " + err.Error())
	}
	viper.Unmarshal(&C)
	C.App.LogStd = viper.GetBool("logstd")

}

func getConfPath() string {
	dir, _ := os.Getwd()

	index := strings.LastIndex(dir, "daimon-server")
	if index < 0 {
		panic("could not run out of project!")
	}
	return filepath.Join(dir[:index], "daimon-server/config")
}
