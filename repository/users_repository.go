package repository

import "github.com/jmoiron/sqlx"

type UsersDB struct {
	db *sqlx.DB
}

func NewUsersDB(db *sqlx.DB) *UsersDB {
	return &UsersDB{db: db}
}

func (u *UsersDB) CreateUser() {

}
