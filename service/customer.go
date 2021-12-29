package service

type CustomerResponse struct {
	CustomerID  int    `json:"customer_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

// ตรงนี้ให้ Define ว่า Business ต้องการอะไรบ้าง
type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
