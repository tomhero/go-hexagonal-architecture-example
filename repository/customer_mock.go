package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

// เอาไว้ mock data จึงไม่ต้องการใช้ database อะไรเลย
func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Tom", City: "Bangkok", ZipCode: "10260", DateOfBirth: "1996-06-16"},
		{CustomerID: 1002, Name: "June", City: "SamutPrakarn", ZipCode: "10130", DateOfBirth: "1996-06-28"},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
