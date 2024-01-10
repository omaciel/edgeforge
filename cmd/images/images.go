package images

import (
	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/spf13/cobra"
)

// var (
// 	imageID int

// 	name         string
// 	version      int
// 	distribution string
// 	arch         string
// 	packages     []string
// 	username     string
// 	sshKey       string
// )

// var DEFAULT_OUTPUT_TYPE = []string{"rhel-edge-installer", "rhel-edge-commit"}

// var cmdImage = &cobra.Command{
// 	Use:   "image",
// 	Short: "Manage your image sets images",
// }

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
