package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	models "github.com/omaciel/edgeforge/pkg/models/images"
	"github.com/spf13/cobra"
)

var cmdImage = &cobra.Command{
	Use:   "image",
	Short: "Manage your images",
}

var cmdCreateImage = &cobra.Command{}

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
		fmt.Println("ImageBuildIsoURL:", response.ImageBuildIsoURL)
		fmt.Println("ImageSet Name:", response.ImageSet.Name)
		fmt.Println("LastImageDetails ID:", response.LastImageDetails.Image.ID)
		fmt.Println("LastImageDetails Version:", response.LastImageDetails.Image.Version)
		fmt.Println("LastImageDetails Description:", response.LastImageDetails.Image.Description)

	},
}

func init() {
	rootCmd.AddCommand(cmdImage)
	cmdImage.AddCommand(cmdCreateImage)
	cmdImage.AddCommand(cmdImageDetails)
	cmdImage.AddCommand(cmdImageSetViewCmd)
}
