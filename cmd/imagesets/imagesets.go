package imagesets

import (
	"github.com/spf13/cobra"
)

type imageSetsCmd struct {
	Cmd *cobra.Command
}

func NewImageSetsCmd() *imageSetsCmd {
	root := &imageSetsCmd{}
	cmd := &cobra.Command{
		Use:   "image-sets",
		Short: "Manage your image sets",
	}

	cmd.AddCommand(
		NewImageSetsListCmd().Cmd,
		NewImageSetsVersionsCmd().Cmd,
	)

	root.Cmd = cmd

	return root
}
