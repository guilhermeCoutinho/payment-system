package dal

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

type DAL struct {
	Account     *Account
	Transaction *Transaction

	db     DB
	config *viper.Viper
}

func NewDAL(
	config *viper.Viper,
	db DB,
) *DAL {
	return &DAL{
		Account:     NewAccount(config, db),
		Transaction: NewTransaction(config, db),
		db:          db,
		config:      config,
	}
}

// DB represents the contract for a Postgres DB
type DB interface {
	//ORM
	//Queryable
	Close() error
	Begin() (*pg.Tx, error)
	Context() context.Context
	WithContext(ctx context.Context) *pg.DB
}

func (dal *DAL) RunInTransaction(ctx context.Context, tx func(txDAL *DAL) error) error {
	//	return dal.db.WithContext(ctx).RunInTransaction(ctx, func(t *pg.Tx) error {
	//		txDAL := NewDAL(dal.config, t)
	return tx(dal)
	//	})
}
