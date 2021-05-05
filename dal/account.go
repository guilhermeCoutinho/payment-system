package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/spf13/viper"
)

type AccountDAL interface {
	Insert(ctx context.Context, account *models.Account) error
	Upsert(ctx context.Context, account *models.Account) error
	Get(ctx context.Context, id uuid.UUID) (*models.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
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

func (a *Account) Insert(ctx context.Context, account *models.Account) error {
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	_, err := a.db.Model(account).Insert()
	return err
}

func (u *Account) Upsert(ctx context.Context, account *models.Account) error {
	account.UpdatedAt = time.Now()
	_, err := u.db.Model(account).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (u *Account) Delete(ctx context.Context, id uuid.UUID) error {
	Account := models.Account{
		ID: id,
	}

	result, err := u.db.Model(&Account).Where("id = ?", id).Delete()
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows")
	}
	return err
}

func (u *Account) Get(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	acc := &models.Account{}

	err := u.db.Model(acc).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return acc, nil
}
