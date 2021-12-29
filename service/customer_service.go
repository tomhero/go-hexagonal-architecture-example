package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

// อะไรที่จะ Public ออกไปให้เอาแค่ที่ต้องการ
func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, error := s.custRepo.GetAll()
	if error != nil {
		log.Println(error)
		return nil, error
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
			return nil, errors.New("customer not found")
		}
		log.Println(err)
		return nil, err
	}
	custResponse := CustomerResponse{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		DateOfBirth: customer.DateOfBirth,
	}
	return &custResponse, nil
}
