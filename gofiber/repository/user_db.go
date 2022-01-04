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

	_, err := r.db.Exec(query, user.username, user.password, user.role, user.customer_id)
	if err != nil {
		return err
	}
	return nil
}

func (r userRespositoryDB) GetByUsername(userName string) (*User, error) {
	user := User{}
	query := "SELECT username, password, role, customer_id FROM users where username=?"

	err := r.db.Get(&user, query, userName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
