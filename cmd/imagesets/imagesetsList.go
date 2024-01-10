package imagesets

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

type imagesetsListCmd struct {
	Cmd    *cobra.Command
	client *clients.APIClient
	opts   imagesetsListOpts
}

type imagesetsListOpts struct {
	imageID int
}

func NewImageSetsListCmd(client *clients.APIClient) *imagesetsListCmd {
	root := &imagesetsListCmd{
		client: client,
	}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all image sets",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var imageSetView types.ImageSetsListResponseAPI

			resp, err := root.client.GetImageSetsList()
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
		},
	}

	root.Cmd = cmd

	return root
}
