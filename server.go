package goapiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iriskin77/goapiserver/handlers"
	"github.com/iriskin77/goapiserver/repository"
	"github.com/iriskin77/goapiserver/service"
	"github.com/iriskin77/goapiserver/store"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	dbConfig     *store.ConfigDB
	serverConfig *ConfigServer
	logger       *logrus.Logger
	router       *mux.Router
}

// Возвращает APiserver со всеми настройками
func NewApiServer(serverConfig *ConfigServer, dbConfig *store.ConfigDB) *APIServer {
	return &APIServer{
		serverConfig: serverConfig,
		dbConfig:     dbConfig,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
	}
}

// Запускает сервер
func (s *APIServer) RunServer() error {

	s.logger.Info("starting API Server")

	db, err := store.NewPostgresDB(*s.dbConfig)

	if err != nil {
		logrus.Fatal("failed to initialize db: %s", err.Error())
	}

	logrus.Info("db has been initialized")

	// Инициализируем репозиторий для работы с БД

	repo := repository.NewRepository(db) // возвращает репозиторий (struct) Repository с методами для БД (CreateUser...)

	service := service.NewService(repo) // возвращает сервис (struct) Service с методами для БД (CreateUser...)

	handlers := handlers.NewHandler(service) // возвращает хэндлеры (struct) Handler

	// Через struct Handler вызываем функции, где регистрируем все хэндлеры, в зависимости от домена

	handlers.RegisterHandlersUsers(s.router)

	return http.ListenAndServe(s.serverConfig.BindAddr, s.router)

}
