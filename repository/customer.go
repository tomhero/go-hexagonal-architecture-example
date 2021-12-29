package repository

/**
ไฟล์นี้คือ Port
*/

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int8   `db:"status"`
}

// Port --> ให้ สร้างเป็น Interface
type CustomerRepository interface {
	GetAll() ([]Customer, error)    // อย่าลืมกำหนดให้มัน return error ออกไปด้วย
	GetById(int) (*Customer, error) // return pointer ของ customer และ error
}
