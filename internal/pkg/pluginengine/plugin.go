package pluginengine

import (
	"github.com/spf13/viper"

	"github.com/merico-dev/stream/internal/pkg/configloader"
)

const DefaultPluginDir = ".devstream"

// DevStreamPlugin is a struct, on which install/reinstall/uninstall interfaces are defined.
type DevStreamPlugin interface {
	// Install will return (true, nil) if there is no error occurred. Otherwise (false, error) will be returned.
	Install(*map[string]interface{}) (bool, error)
	Reinstall(*map[string]interface{}) (bool, error)
	Uninstall(*map[string]interface{}) (bool, error)
	IsHealthy(*map[string]interface{}) (bool, error)
}

// Install loads the plugin and calls the Install method of that plugin.
func Install(tool *configloader.Tool) (bool, error) {
	pluginDir := getPluginDir()
	p, err := loadPlugin(pluginDir, tool)
	if err != nil {
		return false, err
	}
	return p.Install(&tool.Options)
}

// Reinstall loads the plugin and calls the Reinstall method of that plugin.
func Reinstall(tool *configloader.Tool) (bool, error) {
	pluginDir := getPluginDir()
	p, err := loadPlugin(pluginDir, tool)
	if err != nil {
		return false, err
	}
	return p.Reinstall(&tool.Options)
}

// Uninstall loads the plugin and calls the Uninstall method of that plugin.
func Uninstall(tool *configloader.Tool) (bool, error) {
	pluginDir := getPluginDir()
	p, err := loadPlugin(pluginDir, tool)
	if err != nil {
		return false, err
	}
	return p.Uninstall(&tool.Options)
}

func IsHealthy(tool *configloader.Tool) (bool, error) {
	pluginDir := getPluginDir()
	p, err := loadPlugin(pluginDir, tool)
	if err != nil {
		return false, err
	}
	return p.IsHealthy(&tool.Options)
}

func getPluginDir() string {
	var pluginDir string
	if pluginDir = viper.GetString("plugin-dir"); pluginDir == "" {
		pluginDir = DefaultPluginDir
	}
	return pluginDir
}
