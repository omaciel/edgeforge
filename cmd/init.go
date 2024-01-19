package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const defaultForgefile = `# https://github.com/omaciel/edgeforge
baseURL: "https://console.redhat.com/api/edge/v1"
username: ""
password: ""
proxy: ""
verbose: false
`

const defaultForgefileName = ".forge.yaml"

var ErrForgefileAlreadyExists = errors.New("forge: A Forge configuration file already exists")

// InitForgefile Taskfile creates a new Taskfile
func InitForgefile(w io.Writer, dir string) error {
	f := filepath.Join(dir, defaultForgefileName)

	if _, err := os.Stat(f); err == nil {
		return ErrForgefileAlreadyExists
	}

	if err := os.WriteFile(f, []byte(defaultForgefile), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(w, "%s created in the current directory\n", defaultForgefile)
	return nil
}
