package goapiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iriskin77/goapiserver/store"
	"github.com/iriskin77/goapiserver/users"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type APIServer struct {
	serverConfig *ConfigServer
	logger       *logrus.Logger
	router       *mux.Router
}

// Возвращает APiserver со всеми настройками
func NewApiServer(serverConfig *ConfigServer) *APIServer {
	return &APIServer{
		serverConfig: serverConfig,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
	}
}

// Запускает сервер
func (s *APIServer) RunServer() error {

	s.logger.Info("starting API Server")

	// Здесь зарегистрировать все маршруты, инициализируем handlers

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config %s", err.Error())
	}

	// fmt.Println(viper.GetString("db.host"))

	// Инициализруем подключение к БД

	db, err := store.NewPostgresDB(store.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal("failed to initialize db: %s", err.Error())
	}

	logrus.Info("db has been initialized")

	// Инициализируем репозиторий для работы с БД

	repoUsers := users.NewRepository(db) // возвращает интерфейс Repository с методами для БД (CreateUser...)

	usersHander := users.NewHandlerUsers(repoUsers) // Добавляет интерфейс Repository с методами для БД (CreateUser...) в хэндлеры
	usersHander.RegisterHandlersUsers(s.router)

	return http.ListenAndServe(s.serverConfig.BindAddr, s.router)

}

func initConfig() error {
	viper.AddConfigPath("/home/abc/Рабочий стол/goapiserver/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
