package config

import (
	"os"
	"path/filepath"
)

const (
	defaultDir  = ".ein-framework"
	defaultFile = "config.yaml"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	FrameworkRoot string         `yaml:"framework_root"`
	PluginsDir    string         `yaml:"plugins_path"`
	TemplatesDir  string         `yaml:"template_path"`
	Database      DatabaseConfig `yaml:"database"`

	ClientPort     int    `yaml:"client_port"`
	ServerHTTPPort int    `yaml:"server_http_port"`
	Host           string `yaml:"host"`
	SecretToken    string `yaml:"secret_token"`
}

func DefaultPath() string {
	return filepath.Join(defaultDir, defaultFile)
}

func CreateDefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	rootDir := filepath.Join(home, defaultDir)

	return &Config{
		FrameworkRoot: rootDir,
		PluginsDir:    filepath.Join(rootDir, "/plugins"),
		TemplatesDir:  filepath.Join(rootDir, "/templates"),
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "ein",
			Password: "password",
			Name:     "ein",
		},
		ServerHTTPPort: 8081,
		Host:           "localhost",
		SecretToken:    "changeme",
	}
}
