package images

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	imageUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing image",
		Run:   runImageUpdateCmd,
	}
)

func runImageUpdateCmd(cmd *cobra.Command, args []string) {
	client := clients.Get()
	fmt.Println("Image ID:", imageID)

	resp, err := client.GetImageDetails(imageID)
	if err != nil {
		fmt.Println("error fetching image details:", err)
	}

	// Handle the response as needed
	log.Debug("Response Status:", resp.Status())
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("unable to find image with ID:", imageID)
		return
	}

	var response types.Image
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		fmt.Println("error unmarshalling response:", err)
		return
	}

	// Increase the version by one
	response.Version += 1

	if distribution != "" {
		response.Distribution = distribution
	}

	if description != "" {
		response.Description = viper.GetString("description")
	}

	if arch != "" {
		response.Commit.Arch = arch
	}

	if username != "" {
		response.Installer.Username = username
	}

	if sshKey != "" {
		response.Installer.SSHKey = sshKey
	}

	var imageArtifacts = []string{string(flagOutputType)}
	if string(flagOutputType) == string(types.EdgeInstaller) {
		// If building an installer, also explicitly build a commit.
		imageArtifacts = append(imageArtifacts, string(types.EdgeCommit))
	}
	response.OutputTypes = imageArtifacts

	resp, err = client.UpdateImage(imageID, &response)
	if err != nil {
		fmt.Println("error fetching image details:", err)
	}

	// Handle the response as needed
	log.Debug("Response Status:", resp.Status())
}

func init() {
	imageUpdateCmd.Flags().IntVarP(&imageID, "id", "", 0, "Image ID")
	imageUpdateCmd.Flags().StringVarP(&description, "description", "", "", "Distribution")
	imageUpdateCmd.Flags().StringVarP(&distribution, "distribution", "", "", "Distribution")
	imageUpdateCmd.Flags().Var(&flagOutputType, "output-types", `must be one of "rhel-edge-commit", or "rhel-edge-installer"`)
	imageUpdateCmd.Flags().StringVarP(&arch, "arch", "", "", "Architecture")
	imageUpdateCmd.Flags().StringSliceVarP(&packages, "packages", "", nil, "Installed packages")
	imageUpdateCmd.Flags().StringVarP(&username, "ssh-username", "u", "", "Installer username")
	imageUpdateCmd.Flags().StringVarP(&sshKey, "ssh-key", "k", "", "SSH key")
	viper.BindPFlags(imageUpdateCmd.Flags())
	imageUpdateCmd.MarkFlagRequired("id")

	imageCmd.AddCommand(imageUpdateCmd)
}
