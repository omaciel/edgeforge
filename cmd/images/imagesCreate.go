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

var flagOutputType = types.EdgeInstaller

type imageCreateCmd struct {
	Cmd    *cobra.Command
	client *clients.APIClient
	opts   imageCreateOpts
}

type imageCreateOptsFunc func(*imageCreateOpts)

type imageCreateOpts struct {
	name         string
	version      int
	distribution string
	arch         string
	packages     []string
	outputType   string
	username     string
	sshKey       string
}

func WithRHEL8Image(opts *imageCreateOpts) {
	opts.distribution = "rhel-89"
	opts.arch = "x86_64"
}

func WithRHEL9Image(opts *imageCreateOpts) {
	opts.distribution = "rhel-93"
	opts.arch = "x86_64"
}

func WithEdgeInstaler(opts *imageCreateOpts) {
	opts.outputType = string(types.EdgeInstaller)
}

func NewImageCreateCmd(client *clients.APIClient, opts ...imageCreateOptsFunc) *imageCreateCmd {
	root := &imageCreateCmd{
		client: client,
	}

	for _, fn := range opts {
		fn(&root.opts)
	}

	cmd := &cobra.Command{
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

			resp, err := root.client.CreateImage(imagePayload)
			if err != nil {
				log.Fatalf("POST request failed: %v", err)
			}

			log.Debug("Response Status:", resp.Status())

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

	cmd.Flags().StringVarP(&root.opts.name, "name", "n", "", "Image name")
	cmd.Flags().IntVarP(&root.opts.version, "version", "v", 1, "Image version")
	cmd.Flags().StringVarP(&root.opts.distribution, "distribution", "d", "", "Distribution")
	cmd.Flags().Var(&flagOutputType, "output-types", `must be one of "rhel-edge-commit", or "rhel-edge-installer"`)
	cmd.Flags().StringVarP(&root.opts.arch, "arch", "a", "", "Architecture")
	cmd.Flags().StringSliceVarP(&root.opts.packages, "packages", "p", nil, "Installed packages")
	cmd.Flags().StringVarP(&root.opts.username, "ssh-username", "u", "", "Installer username")
	cmd.Flags().StringVarP(&root.opts.sshKey, "ssh-key", "k", "", "SSH key")
	viper.BindPFlags(cmd.Flags())

	root.Cmd = cmd

	return root
}
