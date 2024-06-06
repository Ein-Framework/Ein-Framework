package main

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
	"github.com/google/uuid"
)

type HttpProtocol struct {
	plugins.Plugin
	targetIp   string
	port_range string
	agentId    uuid.UUID
}

func New() plugins.IPlugin {
	return &HttpProtocol{
		Plugin: plugins.Plugin{
			Metadata: plugins.Metadata{
				Version:     "1.0",
				Author:      "Ein-Framework",
				Tags:        []string{"http", "raw-http"},
				ReleaseDate: "2024-06-06",
				SourceLink:  "https://github.com/Ein-Framework/",
				Description: "Scan web applications using http protocol",
			},
			PluginInfo: plugins.PluginInfo{
				Name:       "HttpProtocol",
				Options:    map[string]string{},
				ReturnType: "string",
				Protocol:   "http",
			},
		},
	}
}

func (p HttpProtocol) MetaInfo() *plugins.Metadata {
	return &p.Metadata
}

func (p HttpProtocol) Info() *plugins.PluginInfo {
	return &p.PluginInfo
}

func (p HttpProtocol) Options() map[string]string {
	return p.PluginInfo.Options
}

func (p *HttpProtocol) SetArgs(args map[string]interface{}) error {
	fmt.Println("Setting args")
	return nil
}

func (p *HttpProtocol) Execute(body map[string]interface{}, others ...interface{}) interface{} {
	fmt.Println("Executing")
	fmt.Println(body)
	return nil
}
