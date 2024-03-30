package repository

import (
	"fmt"

	"github.com/iriskin77/goapiserver/models"
	"github.com/jmoiron/sqlx"
)

type UsersDB struct {
	db *sqlx.DB
}

func NewUsersDB(db *sqlx.DB) *UsersDB {
	return &UsersDB{db: db}
}

const (
	usersTable = "users"
)

// Получение пользователя по имени и паролю
func (u *UsersDB) GetUserByPasswordUsername(username, password string) (models.User, error) {

	fmt.Println(username)
	fmt.Println(password)

	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1 AND password_hash = $2", usersTable)
	err := u.db.Get(&user, query, username, password)

	fmt.Println(user)

	return user, err
}

// Создание пользователя
func (u *UsersDB) CreateUser(user *models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, age, password_hash, email) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := u.db.QueryRow(query, user.Name, user.Surname, user.Age, user.Password_hash, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

// Получение пользователя по ID
func (u *UsersDB) GetUserByID(id int) (*models.User, error) {

	var userById models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := u.db.Get(&userById, query, id)

	return &userById, err

}

// Получение списка пользователей
func (u *UsersDB) GetListUsers() ([]models.User, error) {

	var users []models.User

	query := fmt.Sprintf("SELECT name, surname, age, email FROM %s WHERE", usersTable)

	if err := u.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UsersDB) UpdateUserByID(user *models.User) (*models.User, error) {

	userById := &models.User{}

	return userById, nil
}

func (u *UsersDB) DeleteUserByID(id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := u.db.Exec(query, id)

	return err
}
