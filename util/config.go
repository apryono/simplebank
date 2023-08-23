package util

import "github.com/spf13/viper"

// Config store all configuration of the application
// The value are read by the viper from a config file or environment variable
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from the file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // Location of config file
	viper.SetConfigName("app") // to tell viper to look for a config file with the spesific name
	viper.SetConfigType("env") // just make sure the config file has correct format and extention, can use json, xml, dll.

	viper.AutomaticEnv() // to tell viper to automatically override values

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
