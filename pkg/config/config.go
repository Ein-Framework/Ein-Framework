package config

import (
	"os"
	"path/filepath"
)

const (
	defaultDir  = ".ein-framework"
	defaultFile = "config.yaml"
	defaultPath = defaultDir + "/" + defaultFile
)

type Config struct {
	FrameworkRoot string `yaml:"framework_root"`
	PluginsDir    string `yaml:"plugins_path"`
	TemplatesDir  string `yaml:"template_path"`
	ServerDb      string `yaml:"server_db_path"`

	ClientPort     int    `yaml:"client_port"`
	ServerHTTPPort int    `yaml:"server_http_port"`
	Host           string `yaml:"host"`
	SecretToken    string `yaml:"secret_token"`
}

func CreateDefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	rootDir := filepath.Join(home, defaultDir)

	return &Config{
		FrameworkRoot:  rootDir,
		PluginsDir:     filepath.Join(rootDir, "/plugins"),
		TemplatesDir:   filepath.Join(rootDir, "/templates"),
		ServerDb:       filepath.Join(rootDir, "/server.db"),
		ServerHTTPPort: 8081,
		Host:           "localhost",
		SecretToken:    "changeme",
	}
}
