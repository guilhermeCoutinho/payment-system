package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/guilhermeCoutinho/payment-system/models"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
	"github.com/guilhermeCoutinho/payment-system/utils"
)

type GetAccountParams struct {
	ID *uuid.UUID `json:"id"`
}

func (c *GetAccountParams) Validate() error {
	return utils.ValdiateFields(c)
}

func (a *Account) Get(ctx context.Context, args *struct{}, vars *GetAccountParams) (*CreateAccountResponse, *wrapper.HandlerError) {
	err := vars.Validate()
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusBadRequest}
	}

	account, err := a.dal.Account.Get(ctx, *vars.ID)
	if err != nil {
		return nil, &wrapper.HandlerError{Err: err, StatusCode: http.StatusNotFound}
	}

	return &CreateAccountResponse{
		Account: &models.AccountPublicInfo{
			Account:     account,
			CreditLimit: fmt.Sprintf("%.2f", float64(account.CreditLimit)/100.0),
		},
	}, nil
}
