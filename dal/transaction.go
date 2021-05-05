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

type TransactionDAL interface {
	Insert(ctx context.Context, transaction *models.Transaction) error
	Upsert(ctx context.Context, transaction *models.Transaction) error
	Get(ctx context.Context, id uuid.UUID) (*models.Transaction, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Transaction struct {
	config *viper.Viper
	db     *pg.DB
}

func NewTransaction(
	config *viper.Viper,
	db *pg.DB,
) *Transaction {
	return &Transaction{
		config: config,
		db:     db,
	}
}

func (t *Transaction) Insert(ctx context.Context, transaction *models.Transaction) error {
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	_, err := t.db.Model(transaction).Insert()
	return err
}

func (t *Transaction) Upsert(ctx context.Context, transaction *models.Transaction) error {
	transaction.UpdatedAt = time.Now()
	_, err := t.db.Model(transaction).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (t *Transaction) Delete(ctx context.Context, id uuid.UUID) error {
	Transaction := models.Transaction{
		ID: id,
	}

	result, err := t.db.Model(&Transaction).Where("id = ?", id).Delete()
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows")
	}
	return err
}

func (t *Transaction) Get(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	transaction := &models.Transaction{}

	err := t.db.Model(transaction).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
