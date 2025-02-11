package main

import (
	"log"

	"github.com/merico-dev/stream/internal/pkg/plugin/reposcaffolding/github/golang"
)

// NAME is the name of this DevStream plugin.
const NAME = "github-repo-scaffolding-golang"

// Plugin is the type used by DevStream core. It's a string.
type Plugin string

// Install implements the installation of the github-repo-scaffolding-golang.
func (p Plugin) Install(options *map[string]interface{}) (bool, error) {
	return golang.Install(options)
}

// Reinstall implements the reinstallation of the github-repo-scaffolding-golang.
func (p Plugin) Reinstall(options *map[string]interface{}) (bool, error) {
	return golang.Reinstall(options)
}

// Uninstall implements the uninstallation of the github-repo-scaffolding-golang.
func (p Plugin) Uninstall(options *map[string]interface{}) (bool, error) {
	return golang.Uninstall(options)
}

// IsHealthy implements the healthy check of the github-repo-scaffolding-golang.
func (p Plugin) IsHealthy(options *map[string]interface{}) (bool, error) {
	return golang.IsHealthy(options)
}

// DevStreamPlugin is the exported variable used by the DevStream core.
var DevStreamPlugin Plugin

func main() {
	log.Printf("%T: %s is a plugin for DevStream. Use it with DevStream.\n", NAME, DevStreamPlugin)
}
