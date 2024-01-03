package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DEFAULT_OUTPUT_TYPE = []string{"rhel-edge-installer", "rhel-edge-commit"}

var cmdImage = &cobra.Command{
	Use:   "image",
	Short: "Manage your images",
}

var cmdListImageSets = &cobra.Command{
	Use:   "list",
	Short: "Lists all image sets",
	Run: func(cmd *cobra.Command, args []string) {
		var imageSetView types.ImageSetView

		resp, err := client.GetImageSetViews()
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}

		log.Println("Response Status:", resp.Status())

		if err = json.Unmarshal(resp.Body(), &imageSetView); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		if imageSetView.Count > 0 {
			fmt.Printf("%-12s %-32s %-12s\n", "Image ID", "Image Name", "Versions")
			fmt.Printf("%-12s %-32s %-12s\n", "--------", "----------", "--------")
			for _, imgSet := range imageSetView.Data {
				fmt.Printf("%-12d %-32s %-12d\n", imgSet.ID, imgSet.Name, imgSet.Version)
			}
		}
	},
}

var cmdCreateImage = &cobra.Command{
	Use:   "create",
	Short: "Create a new image",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		version := viper.GetInt("version")
		distribution := viper.GetString("distribution")
		outputTypes := viper.GetStringSlice("output-types")
		arch := viper.GetString("arch")
		packages := viper.GetStringSlice("packages")
		username := viper.GetString("ssh-username")
		sshKey := viper.GetString("ssh-key")

		imagePayload := &types.Image{
			Name:         name,
			Version:      version,
			Distribution: distribution,
			OutputTypes:  outputTypes,
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

var cmdImageDetails = &cobra.Command{
	Use:   "details [imageID]",
	Short: "Display details of an image by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the provided image ID
		imageID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error parsing image ID: %s", err)
		}

		resp, err := client.GetImageDetails(imageID)
		if err != nil {
			log.Fatalf("Error fetching image details: %s", err)
		}
		// Handle the response as needed
		log.Println("Response Status:", resp.Status())

		var response types.LastImageDetails
		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("Image ID:", response.Image.ID)
		fmt.Println("Image Name:", response.Image.Name)
		fmt.Println("Image Status:", response.Image.Status)
		fmt.Println("Image Distribution:", response.Image.Distribution)
		fmt.Println("Image Version:", response.Image.Version)
		fmt.Println("Image Description:", response.Image.Description)

	},
}

var cmdImageSetViewCmd = &cobra.Command{
	Use:   "view [imageID]",
	Short: "View details of an image by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the provided image ID
		imageID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error parsing image ID: %s", err)
		}

		resp, err := client.GetImageSetView(imageID)
		if err != nil {
			log.Fatalf("Error fetching image details: %s", err)
		}

		// Handle the response as needed
		log.Println("Response Status:", resp.Status())

		// Process the response body or handle errors
		var response types.ImageSetViewResponseStruct

		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("ImageSet Name:", response.ImageSet.Name)
		fmt.Println("LastImageDetails ID:", response.LastImageDetails.Image.ID)
		fmt.Println("LastImageDetails Version:", response.LastImageDetails.Image.Version)
		fmt.Println("LastImageDetails Description:", response.LastImageDetails.Image.Description)

	},
}

func init() {
	var name string
	var version int
	var distribution string
	var outputTypes []string
	var arch string
	var packages []string
	var username string
	var sshKey string

	cmdCreateImage.Flags().StringVarP(&name, "name", "n", "", "Image name")
	cmdCreateImage.Flags().IntVarP(&version, "version", "v", 0, "Image version")
	cmdCreateImage.Flags().StringVarP(&distribution, "distribution", "d", "", "Distribution")
	cmdCreateImage.Flags().StringSliceVarP(&outputTypes, "output-types", "o", DEFAULT_OUTPUT_TYPE, "Output types")
	cmdCreateImage.Flags().StringVarP(&arch, "arch", "a", "", "Architecture")
	cmdCreateImage.Flags().StringSliceVarP(&packages, "packages", "p", nil, "Installed packages")
	cmdCreateImage.Flags().StringVarP(&username, "ssh-username", "u", "", "Installer username")
	cmdCreateImage.Flags().StringVarP(&sshKey, "ssh-key", "k", "", "SSH key")
	viper.BindPFlags(cmdCreateImage.Flags())

	cmdImage.AddCommand(cmdCreateImage, cmdImageDetails, cmdImageSetViewCmd, cmdListImageSets)

	rootCmd.AddCommand(cmdImage)
}
