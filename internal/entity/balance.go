package entity

type Balance struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}
