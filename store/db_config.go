package store

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfigDB() *ConfigDB {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config %s", err.Error())
	}

	return &ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

}

func initConfig() error {
	viper.AddConfigPath("/home/abc/Рабочий стол/goapiserver/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
