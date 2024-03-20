package users

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *User) (*User, error) {

	if err := r.db.QueryRow(
		"INSERT INTO users (name, surname, age, password_hash, email) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Name,
		user.Surname,
		user.Age,
		user.Password_hash,
		user.Email,
	).Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUserByID(id int) (*User, error) {

	u := &User{}

	if err := r.db.QueryRow("SELECT id, name, surname, email FROM users WHERE id = $1", id).Scan(
		&u.Id,
		&u.Name,
		&u.Surname,
		&u.Email,
	); err != nil {
		logrus.Fatal("unc (r *repository) GetUserByID(id int) (*User, error)")
	}

	return u, nil
}

func (r *repository) GetListUsers() ([]User, error) {
	//q := r.db.QueryRow("SELECT * FROM users")

	//us := make([]User, 0)

	users := make([]User, 0)

	rows, err := r.db.Query("SELECT id, name, surname, age, email FROM users")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var u User

		err = rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Age, &u.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, u)

	}

	return users, nil

	// 2) Instantiate a slice(or struct) which you want to populate, Dummy example.

}
