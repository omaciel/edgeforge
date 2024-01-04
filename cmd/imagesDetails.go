package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var cmdImageDetails = &cobra.Command{
	Use:   "details",
	Short: "Display details of an image by its ID",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := client.GetImageDetails(imageID)
		if err != nil {
			log.Fatalf("Error fetching image details: %s", err)
		}
		// Handle the response as needed
		log.Println("Response Status:", resp.Status())

		var response types.ImageDetail
		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		if response.Image != nil {
			// Access the values in the structured format
			fmt.Println("Image ID:", response.Image.ID)
			fmt.Println("Image Name:", response.Image.Name)
			fmt.Println("Image Status:", response.Image.Status)
			fmt.Println("Image Type:")
			for idx, artifact := range response.Image.OutputTypes {
				fmt.Printf("\t%v - %v\n", idx, artifact)
			}
			fmt.Println("Image Distribution:", response.Image.Distribution)
			fmt.Println("Image Version:", response.Image.Version)
			fmt.Println("Image Description:", response.Image.Description)
			// List any custom packages
			if len(response.Image.Packages) > 0 {
				fmt.Println("Custom Packages:")
				for idx, installedPackage := range response.Image.Packages {
					fmt.Printf("\t%v - %v\n", idx, installedPackage.Name)
				}
			}
		} else {
			fmt.Printf("Image with id '%v' not found.\n", imageID)
		}
	},
}

func init() {
	cmdImageDetails.Flags().IntVarP(&imageID, "id", "", 0, "Image ID")
	cmdImageDetails.MarkFlagRequired("id")
	cmdImage.AddCommand(cmdImageDetails)
}
