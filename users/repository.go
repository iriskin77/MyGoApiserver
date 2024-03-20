package users

type Repository interface {
	CreateUser(*User) (*User, error)
	GetUserByID(id int) (*User, error)
	GetListUsers() ([]User, error)
}
