package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
)

var ErrMissingRequestParam error = fmt.Errorf("request parameter cant be null")

type CreateAccountRequest struct {
	DocumentNumber *string `json:"document_number"`
}

type CreateAccountResponse struct {
	Account *models.Account `json:"account"`
}

func (c *CreateAccountRequest) Validate() error {
	if c.DocumentNumber == nil {
		return ErrMissingRequestParam
	}
	return nil
}

func (a *Account) Post(ctx context.Context, args *CreateAccountRequest, vars *struct{}) (*CreateAccountResponse, *wrapper.HandlerError) {
	err := args.Validate()
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusBadRequest}
	}

	newAccount := a.ToModel(args)
	err = a.dal.Account.Insert(ctx, newAccount)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}

	return &CreateAccountResponse{
		Account: newAccount,
	}, nil
}

func (a *Account) ToModel(args *CreateAccountRequest) *models.Account {
	return &models.Account{
		ID:             uuid.New(),
		DocumentNumber: *args.DocumentNumber,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}
