package imagesets

import (
	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
)

type imageSetsCmd struct {
	Cmd    *cobra.Command
	client *clients.APIClient
}

func NewImageSetsCmd(client *clients.APIClient) *imageSetsCmd {
	root := &imageSetsCmd{
		client: client,
	}
	cmd := &cobra.Command{
		Use:   "image-sets",
		Short: "Manage your image sets",
	}

	cmd.AddCommand(
		NewImageSetsListCmd(root.client).Cmd,
		NewImageSetsVersionsCmd(root.client).Cmd,
	)

	root.Cmd = cmd

	return root
}
