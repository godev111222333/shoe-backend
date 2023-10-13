package store

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbStore struct {
	Db        *gorm.DB
	UserStore *UserStore
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func LoadConfig(path string) (*DbConfig, error) {
	cfg := &DbConfig{}
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

func NewDbStore(cfg *DbConfig) (*DbStore, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &DbStore{
		Db:        db,
		UserStore: NewUserStore(db),
	}, nil
}
