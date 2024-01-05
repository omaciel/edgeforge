package cmd

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

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
		log.Debug("Response Status:", resp.Status())

		var response types.Image
		if err = json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		// Access the values in the structured format
		fmt.Println("Image Name:", response.Name)
		fmt.Println("ID:", response.ID)
		fmt.Println("Distribution:", response.Distribution)
		fmt.Println("Image Version:", response.Version)
		fmt.Println("Image Description:", response.Description)
		fmt.Println("Status:", response.Status)
		fmt.Println("Type:", response.ImageType)
		fmt.Println("OutputTypes:")
		for idx, artifact := range response.OutputTypes {
			fmt.Printf("\t%v - %v\n", idx, artifact)
		}
		// List any custom packages
		if len(response.Packages) > 0 {
			fmt.Println("Custom Packages:")
			for idx, installedPackage := range response.Packages {
				fmt.Printf("\t%v - %v\n", idx, installedPackage.Name)
			}
		}
	},
}

func init() {
	cmdImageDetails.Flags().IntVarP(&imageID, "id", "", 0, "Image ID")
	cmdImageDetails.MarkFlagRequired("id")
	cmdImage.AddCommand(cmdImageDetails)
}
