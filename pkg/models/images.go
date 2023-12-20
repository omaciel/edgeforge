package models

type ImageSet struct {
	ID        int         `json:"ID"`
	CreatedAt string      `json:"CreatedAt"`
	UpdatedAt string      `json:"UpdatedAt"`
	DeletedAt string      `json:"DeletedAt"`
	Name      string      `json:"Name"`
	Version   int         `json:"Version"`
	Account   string      `json:"Account"`
	OrgID     string      `json:"org_id"`
	Images    interface{} `json:"Images"`
}

type ImageSetView struct {
	Count int        `json:"count"`
	Data  []ImageSet `json:"data"`
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

type Commit struct {
	Arch              string             `json:"Arch"`
	InstalledPackages []InstalledPackage `json:"InstalledPackages,omitempty"`
}

type Image struct {
	Name         string     `json:"Name"`
	Distribution string     `json:"Distribution"`
	Description  string     `json:"Description"`
	Version      int        `json:"Version"`
	OutputTypes  []string   `json:"OutputTypes"`
	Commit       *Commit    `json:"Commit"`
	Installer    *Installer `json:"Installer"`
}

type InstalledPackage struct {
	Name      string `json:"name"`
	Arch      string `json:"arch"`
	Release   string `json:"release"`
	Sigmd5    string `json:"sigmd5"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
	Version   string `json:"version"`
	Epoch     string `json:"epoch,omitempty"`
	Commits   []Commit
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
