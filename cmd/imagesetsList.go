package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var cmdImageSetsList = &cobra.Command{
	Use:   "list",
	Short: "Lists all image sets",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var imageSetView types.ImageSetsListResponseAPI

		resp, err := client.GetImageSetsList()
		if err != nil {
			log.Fatalf("request failed: %v", err)
		}

		log.Println("Response Status:", resp.Status())

		if err = json.Unmarshal(resp.Body(), &imageSetView); err != nil {
			log.Fatalln("Error:", err)
			return
		}

		if imageSetView.Count > 0 {
			fmt.Printf("%-12s %-32s %-12s %-12s\n", "Set ID", "Image Name", "Versions", "Distribution")
			fmt.Printf("%-12s %-32s %-12s %-12s\n", "--------", "----------", "--------", "--------")
			for _, imgSet := range imageSetView.Data {
				fmt.Printf("%-12d %-32s %-12d %-12s\n", imgSet.ID, imgSet.Name, imgSet.Version, imgSet.Distribution)
			}
		}
	},
}

func init() {
	cmdImageSets.AddCommand(
		cmdImageSetsList,
	)
}
