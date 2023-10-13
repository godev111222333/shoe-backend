package misc

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	DatabaseConfig *DbConfig  `yaml:"database"`
	APIConfig      *APIConfig `yaml:"api"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type APIConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig(path string) (*GlobalConfig, error) {
	cfg := &GlobalConfig{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	d := yaml.NewDecoder(file)
	if err = d.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
