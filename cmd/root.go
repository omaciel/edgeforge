package cmd

import (
	"log"
	"os"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	client clients.APIClient

	rootCmd = &cobra.Command{
		Use:   "forge",
		Short: "Build personalized Linux images for edge devices with ease.",
		Long:  `Build personalized Linux images for edge devices with ease.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.forge.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		log.Println("User passed configuration file: ", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".forge")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$CWD")
		viper.AddConfigPath(".")
		log.Println("Looking for configuration file.")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Println("configuration file not found: ", err.Error())
		} else {
			// Config file was found but another error was produced
			log.Println("error loading configuration file: ", err.Error())
		}
	}

	log.Println("Using configuration file: ", viper.ConfigFileUsed())

	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()

	settings, err := clients.NewSettings(
		viper.GetString("baseURL"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("proxy"),
	)

	if err != nil {
		log.Fatalf("Error reading settings: %v", err)
	}

	client = *clients.NewAPIClient(settings)
}
