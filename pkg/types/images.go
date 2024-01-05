package types

// ImageDetail return the structure to inform package info to images
type ImageDetail struct {
	Image              *Image `json:"image"`
	AdditionalPackages int    `json:"additional_packages"`
	Packages           int    `json:"packages"`
	UpdateAdded        int    `json:"update_added"`
	UpdateRemoved      int    `json:"update_removed"`
	UpdateUpdated      int    `json:"update_updated"`
}

type LastImageDetails struct {
	Image struct {
		ID           int    `json:"ID"`
		CreatedAt    string `json:"CreatedAt"`
		UpdatedAt    string `json:"UpdatedAt"`
		DeletedAt    string `json:"DeletedAt"`
		Name         string `json:"Name"`
		Account      string `json:"Account"`
		OrgID        string `json:"org_id"`
		Distribution string `json:"Distribution"`
		Description  string `json:"Description"`
		Status       string `json:"Status"`
		Version      int    `json:"Version"`
		ImageType    string `json:"ImageType"`
	} `json:"image"`
}

type ImageSetViewResponseStruct struct {
	ImageBuildIsoURL string           `json:"ImageBuildIsoURL"`
	ImageSet         ImageSet         `json:"ImageSet"`
	LastImageDetails LastImageDetails `json:"LastImageDetails"`
}

type Image struct {
	ID             uint       `json:"CommitID"`
	Name           string     `json:"Name"`
	Distribution   string     `json:"Distribution"`
	Description    string     `json:"Description"`
	Version        int        `json:"Version"`
	Status         string     `json:"Status"`
	ImageType      string     `json:"ImageType"`
	OutputTypes    []string   `json:"OutputTypes"`
	Commit         *Commit    `json:"Commit"`
	Installer      *Installer `json:"Installer"`
	Packages       []Package  `json:"Packages,omitempty"`
	CustomPackages []Package  `json:"CustomPackages,omitempty"`
	RequestID      string     `json:"request_id"`
}

type Installer struct {
	Account          string `json:"Account"`
	OrgID            string `json:"org_id"`
	ImageBuildISOURL string `json:"ImageBuildISOURL"`
	ComposeJobID     string `json:"ComposeJobID"`
	Status           string `json:"Status"`
	Username         string `json:"Username"`
	SSHKey           string `json:"SshKey"`
	Checksum         string `json:"Checksum"`
}

// ImageDetailsResponseAPI is the row returned from v1/image-sets/view/<id>/versions/<id>
type ImageDetailsResponseAPI struct {
	ImageBuildIsoURL string          `json:"ImageBuildIsoURL"`
	ImageSet         ImageSetVersion `json:"ImageSet"`
	ImageDetails     []Image         `json:"image"`
}
