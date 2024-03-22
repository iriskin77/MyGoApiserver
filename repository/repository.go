package repository

import (
	"github.com/jmoiron/sqlx"
)

// Интерфейсы называются в зависимости от участков доменной зоны, за которую они отвечают
type Users interface {
	CreateUser()
}

type Repository struct {
	Users
}

// Конструктор сервисов
// Поскольку репозиторий должен работать с БД, то
func NewRepository(db *sqlx.DB) *Repository {
	// В файле репозитория инициализируем наш репозиторий в конструкторе
	return &Repository{Users: NewUsersDB(db)}
}
