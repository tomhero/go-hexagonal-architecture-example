package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

// อะไรที่จะ Public ออกไปให้เอาแค่ที่ต้องการ
func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	custResponses := []CustomerResponse{}
	// NOTE : Map to CustomerResponse
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID:  customer.CustomerID,
			Name:        customer.Name,
			DateOfBirth: customer.DateOfBirth,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		// NOTE : ตรงนี้ควรจะปั้น Error ระดับ Business เพื่อเอาไว้ Debug ง่าย
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custResponse := CustomerResponse{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		DateOfBirth: customer.DateOfBirth,
	}
	return &custResponse, nil
}
