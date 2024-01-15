package imagesets

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var (
	imageID              int
	imagesetsVersionsCmd = &cobra.Command{
		Use:   "images",
		Short: "Lists all image for an image set",
		Args:  cobra.NoArgs,
		Run:   runImagesetsVersionsCmd,
	}
)

func runImagesetsVersionsCmd(cmd *cobra.Command, args []string) {
	var imageSetView types.ImageSetVersionsResponseAPI

	client := clients.Get()

	resp, err := client.GetImageSetsImages(imageID)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	log.Debug("Response Status:", resp.Status())

	if err = json.Unmarshal(resp.Body(), &imageSetView); err != nil {
		log.Fatalln("Error:", err)
		return
	}

	if imageSetView.Count > 0 {
		fmt.Printf("%-6s %-32s %-6s %-18s %-12s\n", "ID", "Image Name", "Version", "Type", "Status")
		fmt.Printf("%-6s %-32s %-6s %-18s %-12s\n", "------", "----------", "--------", "--------", "--------")
		for _, imgSet := range imageSetView.Data {
			fmt.Printf("%-6d %-32s %-6d %-18s %-12s\n", imgSet.ID, imgSet.Name, imgSet.Version, imgSet.ImageType, imgSet.Status)
		}
	} else {
		fmt.Printf("No images were found for image set with id '%v'.\n", imageID)
	}
}
func init() {
	imagesetsVersionsCmd.Flags().IntVarP(&imageID, "id", "", 0, "Image Set ID")
	imagesetsVersionsCmd.MarkFlagRequired("id")

	imageSetsCmd.AddCommand(imagesetsVersionsCmd)
}
