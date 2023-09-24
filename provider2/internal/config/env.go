package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// init initializes the EnvConfigs variable
func init() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	Secret          string `mapstructure:"SECRET"`
	Provider2Url    string `mapstructure:"PROVIDER2_URL"`
}

// loadEnvVariables loads all environment variables from the provider2.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("provider2")
	// Tell viper the configuration type
	viper.SetConfigType("env")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error while reading the configuration file: ", err)
	}

	// Unmarshal the configuration file into a struct
	var config envConfigs
	err = viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	})
	if err != nil {
		log.Fatal("Error while unmarshalling the configuration file: ", err)
	}

	return &config
}
