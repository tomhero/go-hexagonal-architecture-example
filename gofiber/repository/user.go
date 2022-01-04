package repository

type User struct {
	Username    string `db:"username"`
	Password    string `db:"password"`
	Role        string `db:"role"`
	Customer_id int8   `db:"customer_id"`
}

type UserRepository interface {
	Create(*User) error
	GetByUsername(string) (*User, error)
}
