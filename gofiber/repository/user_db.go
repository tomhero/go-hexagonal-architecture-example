package repository

import "github.com/jmoiron/sqlx"

type userRespositoryDB struct {
	db *sqlx.DB
}

// NOTE : Constructure
func NewUserRepositoryDB(db *sqlx.DB) userRespositoryDB {
	return userRespositoryDB{db: db}
}

// TODO : Implment from UserRepository interface
func (r userRespositoryDB) Create(user User) error {
	return nil
}

func (r userRespositoryDB) GetByUsername(userName string) (*User, error) {
	return nil, nil
}
