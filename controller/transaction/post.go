package transaction

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
	"github.com/guilhermeCoutinho/payment-system/utils"
)

type CreateTransactionRequest struct {
	AccountID     *uuid.UUID `json:"account_id"`
	OperationType *int       `json:"operation_type"`
	Amount        *int64     `json:"amount"`
	// DateInSecondsUTC *int64
}

type CreateTransactionResponse struct {
	Transaction *models.Transaction `json:"transaction"`
}

func (c *CreateTransactionRequest) Validate() error {
	return utils.ValdiateFields(c)
}

func (t *Transaction) Post(ctx context.Context, args *CreateTransactionRequest, vars *struct{}) (*CreateTransactionResponse, *wrapper.HandlerError) {
	err := args.Validate()
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusBadRequest}
	}

	newTransaction := t.ToModel(args)
	err = t.dal.Transaction.Insert(ctx, newTransaction)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}

	return &CreateTransactionResponse{
		Transaction: newTransaction,
	}, nil
}

func (t *Transaction) ToModel(args *CreateTransactionRequest) *models.Transaction {
	return &models.Transaction{
		ID:            uuid.New(),
		AccountID:     *args.AccountID,
		OperationType: models.OperationType(*args.OperationType),
		Amount:        *args.Amount,
		Date:          time.Now().UTC(),

		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
