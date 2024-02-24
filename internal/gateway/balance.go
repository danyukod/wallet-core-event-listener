package gateway

import "github.com/danyukod/wallet-core-event-listener/internal/entity"

type BalanceGateway interface {
	FindByAccountID(accountID string) (*entity.Balance, error)
	CreateBalance(accountID string, balance float64) error
}
