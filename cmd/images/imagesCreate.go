package images

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	imageCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new image",
		Run:   runImageCreateCmd,
	}
)

func runImageCreateCmd(cmd *cobra.Command, args []string) {
	var imageArtifacts = []string{string(flagOutputType)}

	client := clients.Get()

	// If building an installer, also explicitly build a commit.
	if string(flagOutputType) == string(types.EdgeInstaller) {
		imageArtifacts = append(imageArtifacts, string(types.EdgeCommit))
	}
	imagePayload := &types.Image{
		Name:         name,
		Version:      version,
		Distribution: distribution,
		OutputTypes:  imageArtifacts,
		Commit: &types.Commit{
			Arch: arch,
			InstalledPackages: func() []types.InstalledPackage {
				var installedPackages []types.InstalledPackage
				for _, pkg := range packages {
					installedPackages = append(installedPackages, types.InstalledPackage{Name: pkg})
				}
				return installedPackages
			}(),
		},
		Installer: &types.Installer{
			Username: username,
			SSHKey:   sshKey,
		},
	}

	resp, err := client.CreateImage(imagePayload)
	if err != nil {
		log.Println("error creating image:", err)
	}

	log.Debug("Response Status:", resp.Status())

	var response types.Image
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		fmt.Println("error unmarshalling response:", err)
		return
	}

	// Access the values in the structured format
	fmt.Println("Image Created")
	fmt.Println("ID:", response.ID)
	fmt.Println("Name:", response.Name)
	fmt.Println("Distribution:", response.Distribution)
	fmt.Println("Version:", response.Version)
	fmt.Println("Description:", response.Description)
}

func init() {
	imageCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Image name")
	imageCreateCmd.Flags().IntVarP(&version, "version", "", 1, "Image version")
	imageCreateCmd.Flags().StringVarP(&distribution, "distribution", "d", "", "Distribution")
	imageCreateCmd.Flags().Var(&flagOutputType, "output-types", `must be one of "rhel-edge-commit", or "rhel-edge-installer"`)
	imageCreateCmd.Flags().StringVarP(&arch, "arch", "a", "", "Architecture")
	imageCreateCmd.Flags().StringSliceVarP(&packages, "packages", "p", nil, "Installed packages")
	imageCreateCmd.Flags().StringVarP(&username, "ssh-username", "u", "", "Installer username")
	imageCreateCmd.Flags().StringVarP(&sshKey, "ssh-key", "k", "", "SSH key")
	imageCreateCmd.MarkFlagsRequiredTogether("ssh-username", "ssh-key")

	viper.BindPFlags(imageCreateCmd.Flags())

	imageCmd.AddCommand(imageCreateCmd)
}
