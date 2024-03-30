package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iriskin77/goapiserver/models"
	"github.com/sirupsen/logrus"
)

const (
	users = "users/"
	user  = "users/:id"
)

func (h *Handler) RegisterHandlersUsers(router *mux.Router) {
	router.HandleFunc("/createuser", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users", h.GetListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", h.DeleteUserByID).Methods("DELETE")
	router.HandleFunc("/users/{id}", h.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/signin", h.SignIn).Methods("GET")
	router.HandleFunc("/test", h.testAuthorization).Methods("GET")
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	type signInInput struct {
		Name          string `json: "name"`
		Password_hash string `json: "password_hash"`
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzMzU2MTQsImlhdCI6MTcxMTI5MjQxNCwiVXNlcklkIjoyfQ.5K0hTX2YPF05_35nq2CyewVgI6rFP9WuLoCIVAkhN3E
	req := &signInInput{}
	json.NewDecoder(r.Body).Decode(req)

	token, err := h.services.Users.GenerateToken(req.Name, req.Password_hash)

	userToken, err := json.Marshal(token)
	if err != nil {
		logrus.Fatal("CreateUser")
	}

	w.Write([]byte(userToken))

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	newUser := &models.User{}

	//json.NewDecoder(r.Body).Decode(req)
	//fmt.Println(r.Body)
	b, err := json.Marshal(r.Body)
	json.Unmarshal(b, newUser)
	fmt.Println(*newUser)

	user, err := h.services.Users.CreateUser(newUser)

	if err != nil {
		logrus.Fatal("func (h *handlerUsers) CreateUser() didnt work")
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		logrus.Fatal("CreateUser")
	}

	w.Write(resp)

}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	userId, err := strconv.Atoi(id)

	if err != nil {
		// ... handle error
		panic(err)
	}

	user, err := h.services.Users.GetUserByID(userId)
	if err != nil {
		logrus.Fatal("func (h *handlerUsers) GetUserByID(id int) (*User, error)")
	}

	fmt.Println(user)

	resp, err := json.Marshal(user)
	if err != nil {
		logrus.Fatal("GetUserByID")
	}

	w.Write([]byte(resp))

}

func (h *Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	userId, err := strconv.Atoi(id)

	if err != nil {
		// ... handle error
		panic(err)
	}

	err = h.services.Users.DeleteUserByID(userId)

	if err != nil {
		logrus.Fatal("func (h *handlerUsers) DeleteUserByID(id int) (*User, error)")
	}

	resp, err := json.Marshal(http.StatusOK)
	if err != nil {
		logrus.Fatal("DeleteUserByID")
	}

	w.Write(resp)

	//w.Write([]byte("204"))
	//w.Write([]byte(http.StatusOK))

}

func (h *Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {

	type requestUser struct {
		Name          string `json:"name"`
		Surname       string `json:"surname"`
		Age           int    `json:"age"`
		Password_hash string `json:"password_hash"`
		Email         string `json:"email"`
	}

	req := &requestUser{}
	json.NewDecoder(r.Body).Decode(req)

	newUser := &models.User{
		Name:          req.Name,
		Surname:       req.Surname,
		Age:           req.Age,
		Password_hash: req.Password_hash,
		Email:         req.Email,
	}

	user, err := h.services.Users.UpdateUserByID(newUser)

	if err != nil {
		logrus.Fatal("func (h *handlerUsers) CreateUser() didnt work")
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		logrus.Fatal("CreateUser")
	}

	w.Write([]byte(resp))

}

func (h *Handler) GetListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.services.Users.GetListUsers()

	if err != nil {
		logrus.Fatal("func (h *handlerUsers) CreateUser() didnt work")
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		logrus.Fatal("CreateUser")
	}

	w.Write([]byte(resp))

}

func (h *Handler) testAuthorization(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	fmt.Println("Token:", token)

	id, err := h.services.Users.ParseToken(token)

	if err != nil {
		logrus.Fatal("func (h *handlerUsers) testAuthorization() didnt work: %s", err.Error())
		return
	}

	bs := []byte(strconv.Itoa(id))
	w.Write(bs)

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
