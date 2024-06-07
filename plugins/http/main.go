package main

import (
	"fmt"
	"time"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
)

type HttpProtocol struct {
	plugins.Plugin
	alerts []entity.Alert
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
		alerts: make([]entity.Alert, 0),
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

	time.Sleep(time.Second * 30)

	p.alerts = append(p.alerts, entity.Alert{
		Title:       "File Disclosure Found",
		Description: ".env found at http://demo.ein.com",
	})

	p.alerts = append(p.alerts, entity.Alert{
		Title:       "File Disclosure Found",
		Description: ".env.development found at http://demo.ein.com",
	})

	p.alerts = append(p.alerts, entity.Alert{
		Title:       "File Disclosure Found",
		Description: "package.json found at http://demo.ein.com",
	})

	p.alerts = append(p.alerts, entity.Alert{
		Title:       "File Disclosure Found",
		Description: "package-lock.json found at http://demo.ein.com",
	})

	p.alerts = append(p.alerts, entity.Alert{
		Title:       "File Disclosure Found",
		Description: ".git/HEAD found at http://demo.ein.com",
	})

	return entity.TaskExecutionResultType{
		Alerts: p.alerts,
	}
}
