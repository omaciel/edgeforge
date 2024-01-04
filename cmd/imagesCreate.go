package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdCreateImage = &cobra.Command{
	Use:   "create",
	Short: "Create a new image",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		version := viper.GetInt("version")
		distribution := viper.GetString("distribution")
		outputTypes := viper.GetString("output-types")
		arch := viper.GetString("arch")
		packages := viper.GetStringSlice("packages")
		username := viper.GetString("ssh-username")
		sshKey := viper.GetString("ssh-key")

		var imageArtifacts = []string{outputTypes}

		// If building an installer, also explicitly build a commit.
		if outputTypes == string(types.EdgeInstaller) {
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
			log.Fatalf("POST request failed: %v", err)
		}

		log.Println("Response Status:", resp.Status())

		var response types.Image
		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("Image Created")
		fmt.Println("ID:", response.ID)
		fmt.Println("Name:", response.Name)
		fmt.Println("Distribution:", response.Distribution)
		fmt.Println("Version:", response.Version)
		fmt.Println("Description:", response.Description)
	},
}

func init() {
	var flagOutputType = types.EdgeInstaller

	cmdCreateImage.Flags().StringVarP(&name, "name", "n", "", "Image name")
	cmdCreateImage.Flags().IntVarP(&version, "version", "v", 1, "Image version")
	cmdCreateImage.Flags().StringVarP(&distribution, "distribution", "d", "", "Distribution")
	cmdCreateImage.Flags().Var(&flagOutputType, "output-types", `must be one of "rhel-edge-commit", or "rhel-edge-installer"`)
	cmdCreateImage.Flags().StringVarP(&arch, "arch", "a", "", "Architecture")
	cmdCreateImage.Flags().StringSliceVarP(&packages, "packages", "p", nil, "Installed packages")
	cmdCreateImage.Flags().StringVarP(&username, "ssh-username", "u", "", "Installer username")
	cmdCreateImage.Flags().StringVarP(&sshKey, "ssh-key", "k", "", "SSH key")
	viper.BindPFlags(cmdCreateImage.Flags())

	cmdImage.AddCommand(cmdCreateImage)
}
