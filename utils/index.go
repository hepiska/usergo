package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

// //EnvVar function is for read .env file
// func EnvVar(key string) string {
// 	// fmt.Println("EnvVar called")
// 	godotenv.Load(".env")
// 	return os.Getenv(key)
// }

//EnvVar function is for read .env file
func EnvVar(key string) string {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	// var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	// err := viper.Unmarshal(&configuration)

	return viper.GetString(key)

}
