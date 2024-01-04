package cmd

import (
	"github.com/spf13/cobra"
)

var (
	imageID int

	name         string
	version      int
	distribution string
	arch         string
	packages     []string
	username     string
	sshKey       string
)

var DEFAULT_OUTPUT_TYPE = []string{"rhel-edge-installer", "rhel-edge-commit"}

var cmdImage = &cobra.Command{
	Use:   "image",
	Short: "Manage your images",
}

func init() {
	rootCmd.AddCommand(cmdImage)
}
