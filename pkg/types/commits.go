package types

type Commit struct {
	Arch              string             `json:"Arch"`
	InstalledPackages []InstalledPackage `json:"InstalledPackages,omitempty"`
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

// Repo is the delivery mechanism of a Commit over HTTP
type Repo struct {
	URL    string `json:"RepoURL"`
	Status string `json:"RepoStatus"`
}

// Package represents the packages a Commit can have
type Package struct {
	Name string `json:"Name"`
}
