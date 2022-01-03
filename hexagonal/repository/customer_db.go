package repository

/**
ไฟล์นี้คือ Adapter
*/

import "github.com/jmoiron/sqlx"

// Adapter ให้สร้างเป็น type ขึ้นมา (อย่าลืมทำให้เป็น Private ด้วยนะ --> ตัว `c` ตัวเล็ก)
type customerRepositoryDB struct {
	db *sqlx.DB // เราจะทำการ Encapsulate field นี้
}

// คล้ายๆกับ Constructure ใน OOP ยังไงล้า
func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

// Implement ให้เหมือนกับใน Port โดยจะเป็น Receiver function
func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// อย่าลืมว่าต้อง return เป็น Pointer
func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
