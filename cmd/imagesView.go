package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var cmdImageSetView = &cobra.Command{
	Use:   "view",
	Short: "View details of an image by ID",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := client.GetImageSetsView(imageID)
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
	cmdImageSetView.Flags().IntVarP(&imageID, "id", "", 0, "Image ID")
	cmdImageSetView.MarkFlagRequired("id")
	cmdImage.AddCommand(cmdImageSetView)
}
