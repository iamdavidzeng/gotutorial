package main

import (
	"fmt"

	"github.com/iamdavidzeng/gotutorial/03_packages/utils"
)

func main() {
	results := utils.GetSum(1, 2)
	fmt.Println(results)
}
