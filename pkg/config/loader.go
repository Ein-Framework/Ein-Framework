package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	NotFound       string = "[-] error: the configuration file was not found to its specified path. If it's the first time using the framework add `gen-config` or `genc` to generate a config file"
	parsingFailure string = "[-] error: Parsing configuration file has failed"
	UnknownError   string = "[-] error: Unknown error occurred"
)

func GetConfig() (*Config, error) {

	err := GenerateConfigIfNotExists()

	if err != nil {
		log.Fatalln("[-] error getting configuration")
	}

	return parseConfig()

}

func parseConfig() (*Config, error) {

	config, err := ParseConfigFromFile(DefaultPath())
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	return config, nil
}

func ParseConfigFromFile(filePath string) (*Config, error) {

	f, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(NotFound)
		}

		return nil, errors.New(UnknownError)
	}

	var config Config

	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, errors.New(parsingFailure)
	}

	return &config, nil
}
