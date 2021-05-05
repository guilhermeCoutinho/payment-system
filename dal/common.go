package dal

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

type DAL struct {
	Meal AccountDAL
}

func NewDAL(
	config *viper.Viper,
	db *pg.DB,
) *DAL {
	return &DAL{
		Meal: NewAccount(config, db),
	}
}
