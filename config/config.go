package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Sin struct {
		Port   string `yaml:"port"`
		Method string `yaml:"method"`
	} `yaml:"sin"`
	Server struct {
		HealthCheck struct {
			Endpoint      string `yaml:"endpoint"`
			StatusCode    int    `yaml:"statusCode"`
			CheckInterval string `yaml:"checkInterval"`
		} `yaml:"healthCheck"`
		URLs []string `yaml:"urls"`
	} `yaml:"server"`
}

var config *Config

func LoadConfig(file string) (*Config, error) {

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return config, nil

}

func GetConfig() *Config {
	return config
}
