package repository

type User struct {
	username    string `db:"username"`
	password    string `db:"password"`
	role        string `db:"role"`
	customer_id int8   `db:"customer_id"`
}

type UserRepository interface {
	Create(User) error
	GetByUsername(string) (*User, error)
}
