package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySQL struct {
		Username        string `toml:"username"`
		Password        string `toml:"password"`
		Database        string `toml:"database"`
		Host            string `toml:"host"`
		MaxOpenConn     string `toml:"maxopenconn"`
		MaxIdleConn     string `toml:"maxidleconn"`
		ConnMaxLifeTime string `toml:"connmaxlife"`
	} `toml:"mysql"`
}

func LoadConfigs() Config {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("configs")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
	viper.WatchConfig()

	return *config
}
