package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
	"github.com/guilhermeCoutinho/payment-system/utils"
)

type CreateAccountRequest struct {
	DocumentNumber *string `json:"document_number"`
}

type CreateAccountResponse struct {
	Account *models.AccountPublicInfo `json:"account"`
}

func (c *CreateAccountRequest) Validate() error {
	return utils.ValdiateFields(c)
}

func (a *Account) Post(ctx context.Context, args *CreateAccountRequest, vars *struct{}) (*CreateAccountResponse, *wrapper.HandlerError) {
	err := args.Validate()
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusBadRequest}
	}

	newAccount := a.ToModel(args)
	a.enrichCreditLimit(newAccount)

	err = a.dal.Account.Insert(ctx, newAccount)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusInternalServerError}
	}

	return &CreateAccountResponse{
		Account: &models.AccountPublicInfo{
			Account:     newAccount,
			CreditLimit: fmt.Sprintf("%.2f", float64(newAccount.CreditLimit)/100.0),
		},
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

func (a *Account) enrichCreditLimit(account *models.Account) {
	defaultLimit := a.config.GetInt("account.credit.limit")
	account.CreditLimit = defaultLimit * 100
}
