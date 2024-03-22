package service

import "github.com/iriskin77/goapiserver/repository"

type Users interface {
	CreateUser()
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
