package config

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
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
		os.Exit(1)
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
			log.Debug("Configuration file $HOME/.forge.yaml not found:", err.Error())
			fmt.Printf("Configuration file %v not found.\n", viper.ConfigFileUsed())
		} else {
			// Config file was found but another error was produced
			log.Debug("Error loading configuration file:", err.Error())
			fmt.Printf("Error loading configuration file %v.\n", viper.ConfigFileUsed())
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
