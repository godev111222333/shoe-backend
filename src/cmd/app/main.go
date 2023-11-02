package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/godev111222333/shoe-backend/src/handler"
	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/store"
)

func main() {
	cfg, err := misc.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	dbStore, err := store.NewDbStore(cfg.DatabaseConfig)
	if err != nil {
		panic(err)
	}

	otpService := handler.NewOTPService(dbStore, cfg.OTPConfig.Sender, cfg.OTPConfig.Password)

	apiServer := handler.NewAPIServer(cfg.APIConfig, dbStore, otpService)
	go func() {
		if apiServer.Run() != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Press Ctrl+C to exit")
	<-stop
}
