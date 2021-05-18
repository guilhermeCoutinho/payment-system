package transaction

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/dal"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
	"github.com/guilhermeCoutinho/payment-system/utils"
)

type CreateTransactionRequest struct {
	AccountID        *uuid.UUID `json:"account_id" validate:"required"`
	OperationType    *int       `json:"operation_type" validate:"required"`
	Amount           *float64   `json:"amount" validate:"required"`
	DateInSecondsUTC *int64     `json:"date"`
}

type CreateTransactionResponse struct {
	Transaction *models.Transaction `json:"transaction"`
}

func (c *CreateTransactionRequest) Validate() error {
	err := utils.ValdiateFields(c)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Post(ctx context.Context, args *CreateTransactionRequest, vars *struct{}) (*CreateTransactionResponse, *wrapper.HandlerError) {
	err := args.Validate()
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusBadRequest}
	}

	newTransaction := t.ToModel(args)

	account, err := t.dal.Account.Get(ctx, *args.AccountID)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}

	err = t.updateLimit(newTransaction, account)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}

	err = t.dal.RunInTransaction(ctx, func(txDAL *dal.DAL) error {
		err = txDAL.Transaction.Insert(ctx, newTransaction)
		if err != nil {
			return err
		}
		return txDAL.Account.Upsert(ctx, account)
	})

	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}
	return &CreateTransactionResponse{
		Transaction: newTransaction,
	}, nil
}

func (t *Transaction) updateLimit(transaction *models.Transaction, account *models.Account) error {
	isCredit := map[models.OperationType]bool{
		models.Cash:        true,
		models.Installment: true,
		models.Withdraw:    true,
		models.Payment:     false,
	}

	if isCredit[transaction.OperationType] {
		account.CreditLimit -= int(transaction.Amount)
	} else {
		account.CreditLimit += int(transaction.Amount)
	}

	if account.CreditLimit < 0 {
		return fmt.Errorf("operation not allowed")
	}
	return nil
}

func (t *Transaction) ToModel(args *CreateTransactionRequest) *models.Transaction {
	date := time.Now().UTC()
	if args.DateInSecondsUTC != nil {
		date = time.Unix(*args.DateInSecondsUTC, 0)
	}

	return &models.Transaction{
		ID:            uuid.New(),
		AccountID:     *args.AccountID,
		OperationType: models.OperationType(*args.OperationType),
		Amount:        int64(*args.Amount * 100),
		Date:          date,

		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
