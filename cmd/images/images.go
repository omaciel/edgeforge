package images

import (
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var (
	imageID        int
	flagOutputType = types.EdgeInstaller
	name           string
	activationkey  string
	version        int
	distribution   string
	description    string
	arch           string
	packages       []string
	username       string
	sshKey         string
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage your image sets images",
}

func NewImageCmd() *cobra.Command {
	return imageCmd
}
