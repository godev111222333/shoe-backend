package store

import (
	"os"
	"testing"
)

var TestDb *DbStore

func TestMain(m *testing.M) {
	cfg, err := LoadConfig("../../config.yaml")
	if err != nil {
		panic(err)
	}

	initTestDb(cfg)
	code := m.Run()
	os.Exit(code)
}

func initTestDb(cfg *DbConfig) {
	var err error
	TestDb, err = NewDbStore(cfg)
	if err != nil {
		panic(err)
	}
}
