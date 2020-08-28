package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

type DBCfg struct {
	DBURI string `mapstructure:"DB_URI"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config"))
	}

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		envOrRaw := replaceEnvInConfig([]byte(value))
		viper.Set(key, string(envOrRaw))
	}

	var config DBCfg
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to load"))
	}

	fmt.Println(config)
}

func replaceEnvInConfig(body []byte) []byte {
	search := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	replacedBody := search.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := search.ReplaceAllString(string(b), `$1`)
		group2 := search.ReplaceAllString(string(b), `$2`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte(group2)
	})

	return replacedBody
}
