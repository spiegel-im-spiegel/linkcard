package main

import (
	"os"

	"github.com/goark/gocli/rwi"
	"github.com/spiegel-im-spiegel/linkcard/internal/facade"
)

// Version of the application for build info, will be replaced by build script
var Version = "v0.0.0-dev"

// main function of the application
func main() {
	facade.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		Version,
		os.Args[1:],
	).ExitIfNotNormal()
}
