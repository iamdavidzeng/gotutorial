package main

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	DB DBCfg `yaml:"DB"`
}

type DBCfg struct {
	User string `yaml:"USER"`
}

func main() {
	file, err := os.Open("config.yaml")
	defer file.Close()
	if err != nil {
		panic(fmt.Errorf("failed to open file"))
	}

	stream := make([]byte, 1000)

	count, err := file.Read(stream)
	if err != nil {
		panic(fmt.Errorf("i am here"))
	}

	buffer := replaceEnvInConfig(stream[:count])

	var cfg Cfg
	if err := yaml.Unmarshal(buffer, &cfg); err != nil {
		panic(fmt.Errorf("failed to load"))
	}

	fmt.Println(cfg)
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
