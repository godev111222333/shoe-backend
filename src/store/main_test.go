package store

import (
	"os"
	"testing"

	"github.com/godev111222333/shoe-backend/src/misc"
)

var TestDb *DbStore

func TestMain(m *testing.M) {
	cfg, err := misc.LoadConfig("../../config.yaml")
	if err != nil {
		panic(err)
	}

	initTestDb(cfg.DatabaseConfig)
	code := m.Run()
	os.Exit(code)
}

func initTestDb(cfg *misc.DbConfig) {
	var err error
	TestDb, err = NewDbStore(cfg)
	if err != nil {
		panic(err)
	}
}
