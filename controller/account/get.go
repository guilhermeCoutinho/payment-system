package account

import (
	"context"

	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
)

func (m *Account) Get(ctx context.Context, args *struct{}, vars *struct{}) (*struct{}, *wrapper.HandlerError) {
	return nil, nil
}
