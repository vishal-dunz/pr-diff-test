package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println(viper.GetString("test-var"))
}
