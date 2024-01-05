package types

type ImageSetAPI struct {
	Name    string  `json:"name" example:"my-edge-image"` // the image set name
	Version int     `json:"version" example:"1"`          // the image set version
	Images  []Image `json:"Images"`                       // images of image set

}

// ImageSet represents a collection of images
type ImageSet struct {
	ID      uint    `json:"ID"`
	Name    string  `json:"Name"`
	Version int     `json:"Version"`
	Account string  `json:"Account"`
	OrgID   string  `json:"org_id"`
	Images  []Image `json:"Images"`
}

// ImageSetInstallerURL returns Imageset structure with last installer available
type ImageSetInstallerURL struct {
	ImageSetData     ImageSet `json:"image_set"`
	ImageBuildISOURL *string  `json:"image_build_iso_url"`
}

// ImageSetView is the image-set row returned for ui image-sets display
type ImageSetView struct {
	ID               uint     `json:"ID"`
	Name             string   `json:"Name"`
	Version          int      `json:"Version"`
	Distribution     string   `json:"Distribution"`
	OutputTypes      []string `json:"OutputTypes"`
	Status           string   `json:"Status"`
	ImageBuildIsoURL string   `json:"ImageBuildIsoURL"`
	ImageID          uint     `json:"ImageID"`
}

type ImageSetVersion struct {
	ID             uint     `json:"ID"`
	Name           string   `json:"Name"`
	Version        int      `json:"Version"`
	ImageType      string   `json:"ImageType"`
	CommitCheckSum string   `json:"CommitCheckSum"`
	OutputTypes    []string `json:"OutputTypes"`
	Status         string   `json:"Status"`
}

// ImageSetsListResponseAPI is the image-set row returned for ui image-sets display
type ImageSetsListResponseAPI struct {
	Count int            `json:"count" example:"10"` // count of image-sets
	Data  []ImageSetView `json:"data"`               // data of image set view
}

// ImageSetVersionsResponseAPI is the image-set row returned for ui image-sets display
type ImageSetVersionsResponseAPI struct {
	Count int               `json:"count" example:"10"` // count of image-sets
	Data  []ImageSetVersion `json:"data"`               // data of image set view
}
