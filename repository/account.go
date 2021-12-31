package repository

// NOTE : This will create ENUM in Go
// type AccountStatus int8
// const (
// 	Inactive AccountStatus = iota
// 	Active
// )

type Account struct {
	AccountID   int    `db:"account_id"`
	CustomerID  int    `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	Accountype  string `db:"account_type"`
	Amount      int    `db:"amount"`
	Status      int8   `db:"status"`
}

type AccountRepository interface {
	Create(Account) (*Account, error) // NOTE : Create Account
	GetAll() ([]Account, error)       // NOTE : Get All Accounts
}
