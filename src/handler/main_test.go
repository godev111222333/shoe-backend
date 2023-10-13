package handler

import (
	"os"
	"testing"

	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/store"
)

var TestDb *store.DbStore

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
	TestDb, err = store.NewDbStore(cfg)
	if err != nil {
		panic(err)
	}
}
