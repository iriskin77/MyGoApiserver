package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type ConfigDB struct {
// 	Host     string
// 	Port     string
// 	Username string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

// Функция NewPostgresDB возвращает указатель на структуру sqlxDB
func NewPostgresDB(cfg ConfigDB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	// С помощью функции Ping проверяется подключение
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// type Store struct {
// 	db             *sqlx.DB
// }

// func NewStore() *Store {

// 	if err := initConfig(); err != nil {
// 		logrus.Fatalf("error initialization config %s", err.Error())
// 	}

// 	// fmt.Println(viper.GetString("db.host"))

// 	db, err := NewPostgresDB(ConfigDB{
// 		Host:     viper.GetString("db.host"),
// 		Port:     viper.GetString("db.port"),
// 		Username: viper.GetString("db.username"),
// 		Password: viper.GetString("db.password"),
// 		DBName:   viper.GetString("db.dbname"),
// 		SSLMode:  viper.GetString("db.sslmode"),
// 	})

// 	if err != nil {
// 		logrus.Fatal("failed to initialize db: %s", err.Error())
// 	}

// 	return &Store{
// 		db: db,
// 	}
// }

// func initConfig() error {
// 	viper.AddConfigPath("/home/abc/Рабочий стол/goapiserver/configs")
// 	viper.SetConfigName("config")
// 	return viper.ReadInConfig()
// }
