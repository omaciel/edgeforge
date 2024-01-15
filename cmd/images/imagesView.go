package images

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/pkg/clients"
	"github.com/omaciel/edgeforge/pkg/types"
	"github.com/spf13/cobra"
)

var (
	imageID               int
	showCustomPackages    bool
	showInstalledPackages bool
	imageViewCmd          = &cobra.Command{
		Use:   "view",
		Short: "View details of an image by ID",
		Args:  cobra.NoArgs,
		Run:   runImageViewCmd,
	}
)

func runImageViewCmd(cmd *cobra.Command, args []string) {
	client := clients.Get()

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
	// List additional packages
	if showInstalledPackages && len(response.Commit.InstalledPackages) > 0 {
		fmt.Println("Additional Packages:")
		for idx, installedPackage := range response.Commit.InstalledPackages {
			fmt.Printf("\t%v - %v\n", idx, installedPackage.Name)
		}
	}
	// List any custom packages
	if showCustomPackages && len(response.Packages) > 0 {
		fmt.Println("Custom Packages:")
		for idx, installedPackage := range response.Packages {
			fmt.Printf("\t%v - %v\n", idx, installedPackage.Name)
		}
	}
}

func init() {
	imageViewCmd.Flags().IntVarP(&imageID, "id", "", 0, "Image ID")
	imageViewCmd.Flags().BoolVarP(&showCustomPackages, "custom-packages", "", false, "Display additional packages")
	imageViewCmd.Flags().BoolVarP(&showInstalledPackages, "installed-packages", "", false, "Display all installed packages")

	imageViewCmd.MarkFlagRequired("id")

	imageCmd.AddCommand(imageViewCmd)
}
