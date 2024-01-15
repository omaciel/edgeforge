package images

import (
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage your image sets images",
}

func NewImageCmd() *cobra.Command {
	return imageCmd
}
