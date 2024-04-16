package config

import (
	"os"
	"path/filepath"
)

const (
	defaultDir  = ".ein-framework"
	defaultFile = "config.yaml"
	defaultPath = ".ein-framework/config.yaml"
)

type Config struct {
	PluginsDir     string `yaml:"plugins_path"`
	TemplatesDir   string `yaml:"plugins_path"`
	ClientPort     int    `yaml:"client_port"`
	ServergRPCPort int    `yaml:"server_port"`
	ServerHTTPPort int    `yaml:"server_http_port"`
	Host           string `yaml:"host"`
	ServerCert     string `yaml:"server_cert_path"`
	ServerCertKey  string `yaml:"sever_cert_key_path"`
	UseTLS         bool   `yaml:"use_tls"`
	ServerDb       string `yaml:"server_db_path"`
	SecretToken    string `yaml:"secret_token"`
}

func CreateDefaultConfig() *Config {
	home, _ := os.UserHomeDir()

	return &Config{
		PluginsDir:     filepath.Join(home, "/.ein-framework/plugins"),
		TemplatesDir:   filepath.Join(home, "/.ein-framework/templates"),
		ClientPort:     9001,
		ServerHTTPPort: 8081,
		ServergRPCPort: 9002,
		Host:           "localhost",
		ServerCert:     "./cert/server-cert.pem",
		ServerCertKey:  "./cert/server-key.pem",
		UseTLS:         false,
		ServerDb:       "server.db",
	}
}
