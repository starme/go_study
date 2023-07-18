package config

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	config   *viper.Viper
	confOnce sync.Once
)

func makeConfig() *viper.Viper {
	confOnce.Do(func() {
		config = viper.New()

		// read from os env, 配置应用的 env 模式
		mode := os.Getenv("APP_MODE")
		if mode == "" {
			mode = "local"
		}
		config.SetDefault("APP_MODE", mode)
		config.SetDefault("LOG_LEVEL", "info")

		config.AddConfigPath(".")
		config.SetConfigName(mode)
		config.SetConfigType("env")

		if err := config.ReadInConfig(); err != nil {
			panic("read config failed: " + err.Error())
		}

		config = mergeConfig()
	})
	return config
}

func mergeConfig() *viper.Viper {
	for _, file := range Scan("./configs") {
		config.AddConfigPath(file.Path)
		config.SetConfigName(file.Name)
		config.SetConfigType(file.Ext)
		if err := config.MergeInConfig(); err != nil {
			panic("merge config failed: " + err.Error())
		}
	}
	return config
}

func Get(key string) interface{} {
	return makeConfig().Get(key)
}

func GetString(key string) string {
	return makeConfig().GetString(key)
}

func GetInt(key string) int {
	return makeConfig().GetInt(key)
}

func ReadToStruct(s interface{}) error {
	return makeConfig().Unmarshal(s)
}

func Sub(key string) *viper.Viper {
	return makeConfig().Sub(key)
}
