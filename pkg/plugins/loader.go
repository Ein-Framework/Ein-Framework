package plugins

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	Cfg "github.com/Ein-Framework/Ein-Framework/pkg/config"
)

type LoadedPluginInfo struct {
	Plugin  IPlugin
	Channel chan *any
}

type IPluginManager interface {
	ListAllPlugins() ([]string, error)
	ListLoadedPlugins() []string
	LoadAllPlugins() ([]*LoadedPluginInfo, error)
	LoadPlugin(filePath string) (*LoadedPluginInfo, error)
	GetPlugin(filePath string) (*LoadedPluginInfo, error)
}

type PluginManager struct {
	config        *Cfg.Config
	loadedPlugins map[string]*LoadedPluginInfo
	// TaskService   services.ITaskService
}

func CreatePluginManager(cfg *Cfg.Config /*taskService services.ITaskService*/) PluginManager {
	return PluginManager{
		config:        cfg,
		loadedPlugins: make(map[string]*LoadedPluginInfo),
		// TaskService:   taskService,
	}
}

func lookupError(currErr error, errorMsg string) error {
	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: function lookup error in plugin", errorMsg)))
}

// func reflectionError(currErr error, errorMsg string) error {
// 	return errors.Join(currErr, errors.New(fmt.Sprintln("[-] Error: type reflection error in plugin", errorMsg)))
// }

func (manager PluginManager) GetPlugin(filePath string) (*LoadedPluginInfo, error) {
	loadedPlugin, ok := manager.loadedPlugins[filePath]
	if !ok {
		return nil, errors.New("plugin not loaded")
	}
	return loadedPlugin, nil
}

func (manager PluginManager) ListLoadedPlugins() []string {
	keys := make([]string, 0, len(manager.loadedPlugins))
	for k := range manager.loadedPlugins {
		keys = append(keys, k)
	}
	return keys
}

func (manager PluginManager) ListAllPlugins() ([]string, error) {
	var (
		plugins []string
	)

	files, err := os.ReadDir(manager.config.PluginsDir)
	if err != nil {
		// log.Panicln("Error: Cannot load plugins, No directory found.", err)
		return nil, err
	}

	for idx := range files {
		file := files[idx]

		if file.IsDir() || !strings.Contains(file.Name(), ".so") {
			continue
		}

		plugins = append(plugins, file.Name())
	}
	return plugins, nil
}

func (manager *PluginManager) UnloadPlugin(filePath string) error {
	_, ok := manager.loadedPlugins[filePath]
	if !ok {
		return errors.New("plugin not loaded")
	}

	delete(manager.loadedPlugins, filePath)
	return nil
}

func (manager PluginManager) LoadPlugin(filePath string) (*LoadedPluginInfo, error) {

	loadedPlugin, ok := manager.loadedPlugins[filePath]

	if ok {
		fmt.Println("[+] Plugin already loaded:", filePath)
		return loadedPlugin, nil
	}

	fullPath := filepath.Join(manager.config.PluginsDir, filePath)

	fmt.Println("Loading plugin:", filePath)

	err := validateLibraryPath(filePath)
	if err != nil {
		return nil, err
	}

	p, err := plugin.Open(fullPath)
	if err != nil {
		// log.Panicln(err)
		return nil, err
	}

	NewFnSymb, err := p.Lookup("New")
	if err != nil {
		err = lookupError(err, filePath)
		// log.Panicln(err)
		return nil, err
	}

	NewFn, ok := NewFnSymb.(func( /*service services.ITaskService*/ ) IPlugin)
	if !ok {
		return nil, errors.New("new function is not defined")
	}

	pluginInstance := NewFn( /*manager.TaskService*/ )

	validatePlugin(pluginInstance)

	log.Println("[+] Loaded plugin ", pluginInstance.Info().Name)
	manager.loadedPlugins[filePath] = &LoadedPluginInfo{
		Plugin: pluginInstance,
	}
	return manager.loadedPlugins[filePath], nil
}

func (manager PluginManager) LoadAllPlugins() ([]*LoadedPluginInfo, error) {
	var (
		plugins []*LoadedPluginInfo
	)

	files, err := manager.ListAllPlugins()
	if err != nil {
		return nil, err
	}

	for idx := range files {
		file := files[idx]

		loadedPlugin, err := manager.LoadPlugin(file)
		if err != nil {
			fmt.Println("[!] Error loading plugin: ", file)
			return plugins, err
		}

		plugins = append(plugins, loadedPlugin)
	}
	return plugins, nil
}
