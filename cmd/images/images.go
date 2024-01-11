package images

import (
	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
)

type imageCmd struct {
	Cmd    *cobra.Command
	client *clients.APIClient
}

func NewImageCmd(client *clients.APIClient) *imageCmd {
	root := &imageCmd{
		client: client,
	}
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage your image sets images",
	}

	cmd.AddCommand(
		NewImageCreateCmd(root.client).Cmd,
		NewImageViewCmd(root.client).Cmd,
	)

	root.Cmd = cmd

	return root
}
