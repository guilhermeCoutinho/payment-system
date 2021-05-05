package transaction

import (
	"github.com/guilhermeCoutinho/payment-system/dal"
	"github.com/spf13/viper"
)

type Transaction struct {
	dal    *dal.DAL
	config *viper.Viper
}

func NewTransaction(
	dal *dal.DAL,
	config *viper.Viper,
) *Transaction {
	return &Transaction{
		dal:    dal,
		config: config,
	}
}
