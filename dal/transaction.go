package dal

import (
	"context"
	"fmt"
	"time"

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
	db     DB
}

func NewTransaction(
	config *viper.Viper,
	db DB,
) *Transaction {
	return &Transaction{
		config: config,
		db:     db,
	}
}

func (t *Transaction) Insert(ctx context.Context, transaction *models.Transaction) error {
	db := t.db.WithContext(ctx)
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	_, err := db.Model(transaction).Insert()
	return err
}

func (t *Transaction) Upsert(ctx context.Context, transaction *models.Transaction) error {
	db := t.db.WithContext(ctx)
	transaction.UpdatedAt = time.Now()
	_, err := db.Model(transaction).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (t *Transaction) Delete(ctx context.Context, id uuid.UUID) error {
	db := t.db.WithContext(ctx)
	Transaction := models.Transaction{
		ID: id,
	}

	result, err := db.Model(&Transaction).Where("id = ?", id).Delete()
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows")
	}
	return err
}

func (t *Transaction) Get(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	db := t.db.WithContext(ctx)
	transaction := &models.Transaction{}

	err := db.Model(transaction).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
