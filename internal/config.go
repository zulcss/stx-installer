package internal

import (
   "fmt"
   "github.com/spf13/viper"
)

var (
    // Configfile is /etc/stx/installer.yaml per default
    // contains information about the base image
    // and disk which is going to be installed
    ConfigFile string
)

// Reads /etc/stx/installer.yaml configuration
// file
func ReadConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
		viper.ReadInConfig()
		fmt.Printf("Using config file: %s\n", ConfigFile)
	}
}

