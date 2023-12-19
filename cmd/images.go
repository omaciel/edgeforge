package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	models "github.com/omaciel/edgeforge/pkg/models/images"
	"github.com/spf13/cobra"
)

var getImageSetViewCmd = &cobra.Command{
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
		fmt.Println("Response Status:", resp.Status())

		// Process the response body or handle errors
		var response models.ImageSetViewResponseStruct

		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("ImageBuildIsoURL:", response.ImageBuildIsoURL)
		fmt.Println("ImageSet Name:", response.ImageSet.Name)
		fmt.Println("LastImageDetails Description:", response.LastImageDetails.Image.Description)

	},
}

func init() {
	rootCmd.AddCommand(getImageSetViewCmd)
}
