package repository

import "github.com/jmoiron/sqlx"

type userRespositoryDB struct {
	db *sqlx.DB
}

// NOTE : Constructure
func NewUserRepositoryDB(db *sqlx.DB) userRespositoryDB {
	return userRespositoryDB{db: db}
}

func (r userRespositoryDB) Create(user *User) error {
	query := `INSERT INTO users
		(username, password, role, customer_id)
		VALUES(?, ?, ?, ?);`

	_, err := r.db.Exec(query, user.Username, user.Password, user.Role, user.Customer_id)
	if err != nil {
		return err
	}
	return nil
}

func (r userRespositoryDB) GetByUsername(username string) (*User, error) {
	user := User{}
	query := "SELECT username, password, role, customer_id FROM users where username=?"

	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
