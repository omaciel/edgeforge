package types

import "errors"

type OutputType string

const (
	EdgeInstaller OutputType = "rhel-edge-installer"
	EdgeCommit    OutputType = "rhel-edge-commit"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *OutputType) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *OutputType) Set(v string) error {
	switch v {
	case string(EdgeCommit), string(EdgeInstaller):
		*e = OutputType(v)
		return nil
	default:
		return errors.New(`must be one of "rhel-edge-commit", or "rhel-edge-installer"`)
	}
}

// Type is only used in help text
func (e *OutputType) Type() string {
	return "OutputType"
}
