package infrastructure

import (
	"fmt"
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
