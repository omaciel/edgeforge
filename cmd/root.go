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
	"github.com/omaciel/edgeforge/config"
	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

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
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		images.NewImageCmd().Cmd,
		imagesets.NewImageSetsCmd().Cmd,
	)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.forge.yaml)")
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

	config.Init()

	log.Debug("Using configuration file: ", viper.ConfigFileUsed())

	client = *clients.Get()
}
