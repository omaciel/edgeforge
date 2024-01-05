package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	client clients.APIClient

	logLevel string

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

	rootCmd.AddCommand(
		cmdImage,
		cmdImageSets,
	)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.forge.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "INFO", "level of log verbosity")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		log.Debug("User passed configuration file: ", cfgFile)
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

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        time.RFC3339Nano,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		QuoteEmptyFields:       false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	})

	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	log.Debug("Using configuration file: ", viper.ConfigFileUsed())

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
