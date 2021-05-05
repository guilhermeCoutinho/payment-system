package account

import (
	"github.com/guilhermeCoutinho/payment-system/dal"
	"github.com/spf13/viper"
)

type Account struct {
	dal    *dal.DAL
	config *viper.Viper
}

func NewAccount(
	dal *dal.DAL,
	config *viper.Viper,
) *Account {
	return &Account{
		dal:    dal,
		config: config,
	}
}
