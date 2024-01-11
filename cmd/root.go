package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/cmd/images"
	"github.com/omaciel/edgeforge/cmd/imagesets"
	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	config string

	client clients.APIClient

	verbose bool
)

type forgeCmd struct {
	Cmd *cobra.Command
}

func NewForgeCmd() *forgeCmd {
	root := &forgeCmd{}

	cmd := &cobra.Command{
		Use:   "forge",
		Short: "Create personalized Linux images for edge devices with ease.",
		Long:  `Create personalized Linux images for edge devices with ease.`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(
		images.NewImageCmd(&client).Cmd,
		imagesets.NewImageSetsCmd(&client).Cmd,
	)

	cmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.forge.yaml)")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")

	root.Cmd = cmd
	return root
}

// Execute executes the root command.
func Execute() error {
	return NewForgeCmd().Cmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
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
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	if config != "" {
		// Use config file from the flag.
		log.Debug("User passed configuration file: ", config)
		viper.SetConfigFile(config)
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
