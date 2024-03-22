package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	users = "users/"
	user  = "users/:id"
)

func (h *Handler) RegisterHandlersUsers(router *mux.Router) {
	router.HandleFunc("/createuser", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users", h.GetListUsers).Methods("GET")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetListUsers(w http.ResponseWriter, r *http.Request) {

}

// func (h *handlerUsers) CreateUser(w http.ResponseWriter, r *http.Request) {

// 	type request struct {
// 		Name          string `json:"name"`
// 		Surname       string `json:"surname"`
// 		Age           int    `json:"age"`
// 		Password_hash string `json:"password_hash"`
// 		Email         string `json:"email"`
// 	}

// 	req := &request{}
// 	json.NewDecoder(r.Body).Decode(req)

// 	newUser := &User{
// 		Name:          req.Name,
// 		Surname:       req.Surname,
// 		Age:           req.Age,
// 		Password_hash: req.Password_hash,
// 		Email:         req.Email,
// 	}

// 	user, err := h.repository.CreateUser(newUser)

// 	if err != nil {
// 		logrus.Fatal("func (h *handlerUsers) CreateUser() didnt work")
// 		return
// 	}

// 	resp, err := json.Marshal(user)
// 	if err != nil {
// 		logrus.Fatal("CreateUser")
// 	}

// 	w.Write([]byte(resp))

// }

// func (h *handlerUsers) GetUserByID(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	userId, err := strconv.Atoi(id)

// 	if err != nil {
// 		// ... handle error
// 		panic(err)
// 	}

// 	user, err := h.repository.GetUserByID(userId)
// 	if err != nil {
// 		logrus.Fatal("func (h *handlerUsers) GetUserByID(id int) (*User, error)")
// 	}

// 	fmt.Println(user)

// 	resp, err := json.Marshal(user)
// 	if err != nil {
// 		logrus.Fatal("CreateUser")
// 	}

// 	w.Write([]byte(resp))

// }

// func (h *handlerUsers) GetListUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.repository.GetListUsers()
// 	if err != nil {
// 		logrus.Fatal("func (h *handlerUsers) GetUserByID(id int) (*User, error)")
// 	}

// 	resp, err := json.Marshal(users)
// 	if err != nil {
// 		logrus.Fatal("CreateUser")
// 	}

// 	w.Write([]byte(resp))
// }
