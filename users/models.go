package users

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Age           int    `json:"age"`
	Password_hash string `json:"password_hash"`
	Email         string `json:"email"`
	Is_admin      bool   `json:"is_admin"`
}
