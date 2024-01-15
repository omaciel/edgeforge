package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/omaciel/edgeforge/cmd"
)

func main() {

	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.Debug(err)
	}
}
