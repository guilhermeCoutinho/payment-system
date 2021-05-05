package dal

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

type DAL struct {
	Account     AccountDAL
	Transaction TransactionDAL
}

func NewDAL(
	config *viper.Viper,
	db *pg.DB,
) *DAL {
	return &DAL{
		Account:     NewAccount(config, db),
		Transaction: NewTransaction(config, db),
	}
}
