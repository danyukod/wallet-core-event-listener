package create_balance

import "github.com/danyukod/wallet-core-event-listener/internal/gateway"

type CreateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type CreateBalanceUseCase struct {
	gateway.BalanceGateway
}

func NewCreateBalanceUseCase(balanceGateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (c CreateBalanceUseCase) Execute(dto CreateBalanceInputDTO) error {
	err := c.BalanceGateway.CreateBalance(dto.AccountID, dto.Balance)
	if err != nil {
		return err
	}

	return nil
}
