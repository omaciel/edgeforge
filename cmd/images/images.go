package images

import (
	"github.com/spf13/cobra"
)

type imageCmd struct {
	Cmd *cobra.Command
}

func NewImageCmd() *imageCmd {
	root := &imageCmd{}
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage your image sets images",
	}

	cmd.AddCommand(
		NewImageCreateCmd().Cmd,
		NewImageViewCmd().Cmd,
	)

	root.Cmd = cmd

	return root
}
