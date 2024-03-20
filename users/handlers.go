package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	users = "users/"
	user  = "users/:id"
)

type handlerUsers struct {
	repository Repository
}

func NewHandlerUsers(repository Repository) *handlerUsers {
	return &handlerUsers{repository: repository}
}

func (h *handlerUsers) RegisterHandlersUsers(router *mux.Router) {
	router.HandleFunc("/createuser", h.CreateUser)
}

func (h *handlerUsers) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := h.repository.CreateUser()
	w.Write([]byte(res))
}

// 	type request struct {
// 		Name          string `json:"name"`
// 		Surname       string `json:"surname"`
// 		Age           int    `json:"age"`
// 		Password_hash string `json:"password_hash"`
// 		Email         string `json:"email"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			s.error(w, r, http.StatusBadRequest, err)
// 		}

// 		u := &User{
// 			Name:          req.Name,
// 			Surname:       req.Surname,
// 			Age:           req.Age,
// 			Password_hash: req.Password_hash,
// 			Email:         req.Email,
// 		}
// 		// if err :=
// 	}

// }

// func (h *handlerUsers) GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("ListUsers"))
// }

// func (h *handlerUsers) GetListUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("ListUsers"))
// }

// func (h *handlerUsers) error(w.ResponseWriter, r *http.Request, code int, err error) {
// 	s.respond(w, r, code, map[string]string{"error": err.Error()})

// }

// func (h *handlerUsers) respond(w.ResponseWriter, r *http.Request, code int, data interface{}) {
// 	w.WrtieHead(code)
// 	if data != nil {
// 		json.NewEncoder(w).Encode(data)
// 	}

// }
