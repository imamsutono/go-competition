package initializers

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(filePath string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("using config file:", viper.ConfigFileUsed())
}
