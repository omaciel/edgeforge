package config

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Config *config
	lock   = &sync.Mutex{}
)

type config struct {
	BaseURL  string `yaml:"baseurl"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	ProxyUrl string `yaml:"proxy"`
}

func Init() {
	newConfig, err := newConfiguration()
	if err != nil {
		return
	}
	Config = newConfig
}
func Get() *config {
	if Config == nil {
		lock.Lock()
		defer lock.Unlock()
		if Config == nil {
			Config, _ = newConfiguration()
		}
	}
	return Config
}

func newConfiguration() (*config, error) {
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Error("configuration file not found: ", err.Error())
		} else {
			// Config file was found but another error was produced
			log.Error("error loading configuration file: ", err.Error())
		}
		return nil, err
	}

	forgeConfig := &config{
		BaseURL:  viper.GetString("baseURL"),
		Username: viper.GetString("userName"),
		Password: viper.GetString("password"),
		ProxyUrl: viper.GetString("proxy"),
	}
	return forgeConfig, nil
}

func readConfigurationFile(configFile string) {
	if configFile != "" {
		// Use config file from the flag.
		log.Debug("User passed configuration file: ", configFile)
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".forge")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$CWD")
		viper.AddConfigPath(".")
		log.Debug("Looking for configuration file.")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Debug("configuration file not found: ", err.Error())
		} else {
			// Config file was found but another error was produced
			log.Debug("error loading configuration file: ", err.Error())
		}
	}

}
