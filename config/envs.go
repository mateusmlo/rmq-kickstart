package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

// GetEnvs read environment vars from .env file in root
func GetEnvs() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	viper.SetConfigFile(path.Join(dir, ".env"))
	viper.ReadInConfig()
}
