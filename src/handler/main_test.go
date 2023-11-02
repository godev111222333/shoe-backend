package handler

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/store"
)

var (
	OtpService    *OTPService
	TestDb        *store.DbStore
	TestAPIServer *APIServer
)

const (
	ResetDB = false
)

func TestMain(m *testing.M) {
	cfg, err := misc.LoadConfig("../../config.yaml")
	if err != nil {
		panic(err)
	}

	if ResetDB {
		if err := ResetDb(cfg.DatabaseConfig); err != nil {
			panic(err)
		}
	}

	initTestDb(cfg.DatabaseConfig)
	initServices(cfg)
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

func initServices(cfg *misc.GlobalConfig) {
	OtpService = NewOTPService(TestDb, cfg.OTPConfig.Sender, cfg.OTPConfig.Password)
	TestAPIServer = NewAPIServer(cfg.APIConfig, TestDb, OtpService)
}

func ResetDb(cfg *misc.DbConfig) error {
	dbString := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
	downCmd := exec.Command("migrate", "-path", "../../migration", "-database", dbString, "-verbose", "down")
	downCmd.Stdout = os.Stdout
	downCmd.Stderr = os.Stderr
	downStdIn, err := downCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if err := downCmd.Start(); err != nil {
		return err
	}

	// send "y" cmd
	if _, err := io.WriteString(downStdIn, "y\n"); err != nil {
		return err
	}

	if err := downCmd.Wait(); err != nil {
		return err
	}

	upCmd := exec.Command("migrate", "-path", "../../migration", "-database", dbString, "-verbose", "up")
	if err := upCmd.Run(); err != nil {
		return err
	}

	return nil
}
