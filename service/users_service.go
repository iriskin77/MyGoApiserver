package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iriskin77/goapiserver/models"
	"github.com/iriskin77/goapiserver/repository"
	"github.com/sirupsen/logrus"
)

const (
	tokenExpired = 12 * time.Hour
	signingKey   = "ggrqqweqwe#1232wefjgkd"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: id`
}

type ServiceUsers struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	repo repository.Users
}

func NewUsersService(repo repository.Users) *ServiceUsers {
	// Конструктор: принимает репозиторий, возваращает сервис с репозиторием
	return &ServiceUsers{repo: repo}
}

func (s *ServiceUsers) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUserByPasswordUsername(username, password)

	if err != nil {
		logrus.Fatal("Error: func (s *ServiceUsers) GenerateToken(): %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpired).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})

	return token.SignedString([]byte(signingKey))
}

// Функция для парсинга токена
func (s *ServiceUsers) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *ServiceUsers) CreateUser(user *models.User) (int, error) {
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		logrus.Fatal("Error: func (s *ServiceUsers) CreateUser(user *models.User)")
	}
	return newUser, nil

}

func (s *ServiceUsers) GetUserByID(id int) (*models.User, error) {
	userById, err := s.repo.GetUserByID(id)
	if err != nil {
		logrus.Fatal("Error: func (s *ServiceUsers) GetUserByID(id int)")
	}

	return userById, nil
}

func (s *ServiceUsers) GetListUsers() ([]models.User, error) {
	users, err := s.repo.GetListUsers()

	if err != nil {
		logrus.Fatal("Error: func (s *ServiceUsers) GetUserByID(id int)")
	}

	return users, nil
}

func (s *ServiceUsers) UpdateUserByID(user *models.User) (*models.User, error) {
	updatedUserById, err := s.repo.UpdateUserByID(user)

	if err != nil {
		logrus.Fatal("Error: func (s *ServiceUsers) GetUserByID(id int)")
	}

	return updatedUserById, nil
}

func (s *ServiceUsers) DeleteUserByID(id int) error {

	err := s.repo.DeleteUserByID(id)

	return err
}

// func (r *repository) CreateUser(user *User) (*User, error) {

// 	if err := r.db.QueryRow(
// 		"INSERT INTO users (name, surname, age, password_hash, email) VALUES ($1, $2, $3, $4, $5) RETURNING id",
// 		user.Name,
// 		user.Surname,
// 		user.Age,
// 		user.Password_hash,
// 		user.Email,
// 	).Scan(&user.Id); err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (r *repository) GetUserByID(id int) (*User, error) {

// 	u := &User{}

// 	if err := r.db.QueryRow("SELECT id, name, surname, email FROM users WHERE id = $1", id).Scan(
// 		&u.Id,
// 		&u.Name,
// 		&u.Surname,
// 		&u.Email,
// 	); err != nil {
// 		logrus.Fatal("unc (r *repository) GetUserByID(id int) (*User, error)")
// 	}

// 	return u, nil
// }

// func (r *repository) GetListUsers() ([]User, error) {
// 	//q := r.db.QueryRow("SELECT * FROM users")

// 	//us := make([]User, 0)

// 	users := make([]User, 0)

// 	rows, err := r.db.Query("SELECT id, name, surname, age, email FROM users")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for rows.Next() {
// 		var u User

// 		err = rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Age, &u.Email)
// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, u)

// 	}

// 	return users, nil

// }
