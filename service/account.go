package service

// NOTE : Request Body
type NewAccountRequest struct {
	// NOTE : ที่ Comment ไว้คือ Business data ที่ไม่ได้ใช้จาก request body
	// AccountID   int    `json:"account_id"`
	// CustomerID  int    `json:"customer_id"`
	// OpeningDate string `json:"opening_date"`
	Accountype string `json:"account_type"`
	Amount     int    `json:"amount"`
	// Status      int8   `json:"status"`
}

// NOTE : Response Body
type AccountResponse struct {
	AccountID int `json:"account_id"`
	// CustomerID  int    `json:"customer_id"` // NOTE : ไม่ได้ใช้เนื่องจาก Request มันมีการ Search มาจากอันนี้อยู่แล้ว
	OpeningDate string `json:"opening_date"`
	Accountype  string `json:"account_type"`
	Amount      int    `json:"amount"`
	Status      int8   `json:"status"`
}

type AccountService interface {
	NewAcccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(int) ([]AccountResponse, error)
}
