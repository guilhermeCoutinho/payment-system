package dal

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

type AccountDAL interface {
}

type Account struct {
	config *viper.Viper
	db     *pg.DB
}

func NewAccount(
	config *viper.Viper,
	db *pg.DB,
) *Account {
	return &Account{
		config: config,
		db:     db,
	}
}
