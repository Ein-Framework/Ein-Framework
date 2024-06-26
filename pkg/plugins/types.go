package plugins

import (
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
	UnloadPlugin(filePath string) error
	GetPlugin(filePath string) (*LoadedPluginInfo, error)
	GetPluginByProtocol(protocol string) (*LoadedPluginInfo, error)
}

type PluginManager struct {
	config           *Cfg.Config
	loadedPlugins    map[string]*LoadedPluginInfo
	protocolToPlugin map[string]*LoadedPluginInfo
}
