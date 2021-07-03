package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

type DatabaseCfg struct {
	Address string `mapstructure:"DB_URL"`
}

var config DatabaseCfg

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config"))
	}

	for _, key := range viper.AllKeys() {
		viper.Set(key, replace(viper.GetString(key)))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to load"))
	}

	fmt.Println(config)
}

func replace(s string) string {
	compiler := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	value := compiler.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		match := compiler.FindStringSubmatch(string(b))
		fmt.Println(match)

		envValue := os.Getenv(match[1])
		defaultValue := match[2]

		if len([]byte(envValue)) > 0 {
			return []byte(envValue)
		}
		return []byte(defaultValue)
	})

	return string(value)
}
