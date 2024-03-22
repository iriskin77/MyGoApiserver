package service

import "github.com/iriskin77/goapiserver/repository"

// type Service struct {
// 	repository *repository.Repository
// }

// func NewUsersService(repo *repository.Repository) *Service {
// 	return &Service{repository: repo}
// }

// func (s *Service) CreateUser() {

// }

type ServiceUsers struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	repo repository.Users
}

func NewUsersService(repo repository.Users) *ServiceUsers {
	// Конструктор: принимает репозиторий, возваращает сервис с репозиторием
	return &ServiceUsers{repo: repo}
}

func (s *ServiceUsers) CreateUser() {

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
