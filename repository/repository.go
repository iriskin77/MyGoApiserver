package repository

import (
	"github.com/iriskin77/goapiserver/models"
	"github.com/jmoiron/sqlx"
)

// Интерфейсы называются в зависимости от участков доменной зоны, за которую они отвечают
type Users interface {
	GetListUsers() ([]models.User, error)
	UpdateUserByID(user *models.User) (*models.User, error)
	DeleteUserByID(id int) error
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) (int, error)
	GetUserByPasswordUsername(username, password string) (models.User, error)
	//GenerateToken(username, passwrpd string) (int, error)
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
