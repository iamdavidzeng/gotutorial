package main

import (
	"fmt"

	"github.com/iamdavidzeng/gonameko"
)

func main() {
	namekoclient := gonameko.Client{
		RabbitHostname: "localhost",
		RabbitUser:     "guest",
		RabbitPass:     "guest",
		RabbitPort:     5672,
		ContentType:    "application/xjson",
	}
	namekoclient.Setup()

	response, err := namekoclient.Call(
		gonameko.RPCRequestParam{
			Service:  "properties",
			Function: "health_check",
			Payload: gonameko.RPCPayload{
				Args:   []string{},
				Kwargs: map[string]string{},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
