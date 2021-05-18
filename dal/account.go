package dal

import (
	"context"
	"fmt"
	"time"

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
	db     DB
}

func NewAccount(
	config *viper.Viper,
	db DB,
) *Account {
	return &Account{
		config: config,
		db:     db,
	}
}

func (a *Account) Insert(ctx context.Context, account *models.Account) error {
	db := a.db.WithContext(ctx)
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	_, err := db.Model(account).Insert()
	return err
}

func (u *Account) Upsert(ctx context.Context, account *models.Account) error {
	db := u.db.WithContext(ctx)
	account.UpdatedAt = time.Now()
	_, err := db.Model(account).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (u *Account) Delete(ctx context.Context, id uuid.UUID) error {
	db := u.db.WithContext(ctx)
	Account := models.Account{
		ID: id,
	}

	result, err := db.Model(&Account).Where("id = ?", id).Delete()
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows")
	}
	return err
}

func (u *Account) Get(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	db := u.db.WithContext(ctx)
	acc := &models.Account{}

	err := db.Model(acc).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return acc, nil
}
