package imagesets

import (
	"github.com/spf13/cobra"
)

var imageSetsCmd = &cobra.Command{
	Use:   "image-sets",
	Short: "Manage your image sets",
}

func NewImageSetsCmd() *cobra.Command {
	return imageSetsCmd
}
