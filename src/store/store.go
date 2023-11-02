package store

import (
	"fmt"

	"github.com/godev111222333/shoe-backend/src/misc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbStore struct {
	Db        *gorm.DB
	UserStore *UserStore
	OTPStore  *OTPStore
}

func NewDbStore(cfg *misc.DbConfig) (*DbStore, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &DbStore{
		Db:        db,
		UserStore: NewUserStore(db),
		OTPStore:  NewOTPStore(db),
	}, nil
}
