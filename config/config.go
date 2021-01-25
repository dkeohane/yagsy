package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Exported Identifiers must use uppercase first letters
type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Driver   string `yaml:"driver"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Auth struct {
		JwtSecret string `yaml:"jwt_secret"`
	} `yaml:"auth"`
}

func (conf *Config) LoadConfig(fileName string, loadRelativePath bool) error {

	if !loadRelativePath {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error reading pwd: %s\n", err)
			return err
		}
		fileName = pwd + fileName
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return err
	}

	return nil
}
