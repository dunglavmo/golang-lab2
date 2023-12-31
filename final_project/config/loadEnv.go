package config

import (
	"log"

	"github.com/joho/godotenv"
)

// import (
// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	DBHost         string `mapstructure:"POSTGRES_HOST"`
// 	DBUserName     string `mapstructure:"POSTGRES_USER"`
// 	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
// 	DBName         string `mapstructure:"POSTGRES_DB"`
// 	DBPort         string `mapstructure:"POSTGRES_PORT"`
// 	ServerPort     string `mapstructure:"PORT"`

// 	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
// }

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("dev")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

//		err = viper.Unmarshal(&config)
//		return
//	}
func LoadEnv() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
