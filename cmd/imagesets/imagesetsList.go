package imagesets

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var imagesetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all image sets",
	Args:  cobra.NoArgs,
	Run:   runImageListCmd,
}

func runImageListCmd(cmd *cobra.Command, args []string) {
	var imageSetView types.ImageSetsListResponseAPI

	client := clients.Get()

	resp, err := client.GetImageSetsList()
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	log.Debug("Response Status:", resp.Status())

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
	} else {
		fmt.Println("No image sets were found.")
	}
}

func init() {
	imageSetsCmd.AddCommand(imagesetsListCmd)
}
