package service

import (
	"github.com/iriskin77/goapiserver/models"
	"github.com/iriskin77/goapiserver/repository"
)

type Users interface {
	GetListUsers() ([]models.User, error)
	UpdateUserByID(user *models.User) (*models.User, error)
	DeleteUserByID(id int) error
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) (int, error)
	GenerateToken(username, passwrpd string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Users
}

// Конструктор сервисов. Сервисы будут передавать данные из хэндлера ниже, на уровень репозитория, поэтому нужен указатель
// на структуру репозитория (репозиторий коннектиться к БД)

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repository.Users),
	}
}
