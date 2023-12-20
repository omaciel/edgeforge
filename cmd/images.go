package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	models "github.com/omaciel/edgeforge/pkg/models/images"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DEFAULT_OUTPUT_TYPE = []string{"rhel-edge-installer", "rhel-edge-commit"}

var cmdImage = &cobra.Command{
	Use:   "image",
	Short: "Manage your images",
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

		imagePayload := &models.Image{
			Name:         name,
			Version:      version,
			Distribution: distribution,
			OutputTypes:  outputTypes,
			Commit: &models.Commit{
				Arch: arch,
				InstalledPackages: func() []models.InstalledPackage {
					var installedPackages []models.InstalledPackage
					for _, pkg := range packages {
						installedPackages = append(installedPackages, models.InstalledPackage{Name: pkg})
					}
					return installedPackages
				}(),
			},
			Installer: &models.Installer{
				Username: username,
				SSHKey:   sshKey,
			},
		}

		resp, err := client.CreateImage(imagePayload)
		if err != nil {
			log.Fatalf("POST request failed: %v", err)
		}

		log.Println("Response Status:", resp.Status())
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

		var response models.LastImageDetails
		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("ImageDetails ID:", response.Image.ID)
		fmt.Println("ImageDetails Name:", response.Image.Name)
		fmt.Println("ImageDetails Status:", response.Image.Status)
		fmt.Println("ImageDetails Distribution:", response.Image.Distribution)
		fmt.Println("ImageDetails Version:", response.Image.Version)
		fmt.Println("ImageDetails Description:", response.Image.Description)

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
		var response models.ImageSetViewResponseStruct

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

	cmdImage.AddCommand(cmdCreateImage, cmdImageDetails, cmdImageSetViewCmd)

	rootCmd.AddCommand(cmdImage)
}
