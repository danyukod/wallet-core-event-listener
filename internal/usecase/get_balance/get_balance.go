package get_balance

import "github.com/danyukod/wallet-core-event-listener/internal/gateway"

type GetBalanceInputDTO struct {
	AccountID string `json:"account_id"`
}

type GetBalanceOutputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type GetBalanceUseCase struct {
	gateway.BalanceGateway
}

func NewGetBalanceUseCase(balanceGateway gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (c GetBalanceUseCase) Execute(dto GetBalanceInputDTO) (*GetBalanceOutputDTO, error) {
	account, err := c.BalanceGateway.FindByAccountID(dto.AccountID)
	if err != nil {
		return nil, err
	}

	return &GetBalanceOutputDTO{
		AccountID: account.ID,
		Balance:   account.Balance,
	}, nil
}
