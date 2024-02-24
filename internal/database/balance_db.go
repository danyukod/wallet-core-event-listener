package database

import (
	"database/sql"
	"github.com/danyukod/wallet-core-event-listener/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) FindByAccountID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := b.DB.Prepare("SELECT id, account_id, balance FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountID)
	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Balance)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (b *BalanceDB) CreateBalance(accountID string, balance float64) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (account_id, balance) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(accountID, balance)
	if err != nil {
		return err
	}
	return nil
}
