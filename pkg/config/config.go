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
	ServerCert    string `yaml:"server_cert_path"`
	ServerCertKey string `yaml:"sever_cert_key_path"`
	ServerDb      string `yaml:"server_db_path"`

	ClientPort     int    `yaml:"client_port"`
	ServerHTTPPort int    `yaml:"server_http_port"`
	Host           string `yaml:"host"`
	UseTLS         bool   `yaml:"use_tls"`
	SecretToken    string `yaml:"secret_token"`
}

func CreateDefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	rootDir := filepath.Join(home, defaultDir)

	return &Config{
		FrameworkRoot: rootDir,
		PluginsDir:    filepath.Join(rootDir, "/plugins"),
		TemplatesDir:  filepath.Join(rootDir, "/templates"),
		ServerCert:    filepath.Join(rootDir, "/server-cert.pem"),
		ServerCertKey: filepath.Join(rootDir, "/server-key.pem"),
		ServerDb:      filepath.Join(rootDir, "/server.db"),

		ClientPort:     9001,
		ServerHTTPPort: 8081,
		Host:           "localhost",
		UseTLS:         false,
		SecretToken:    "changeme",
	}
}
