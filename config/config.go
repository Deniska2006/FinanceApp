package config

import (
	"github.com/upper/db/v4/adapter/postgresql"
)

var Settings = postgresql.ConnectionURL{
	Database: "FinDB",     // Назва бази
	Host:     "localhost", // Сервер (або IP)
	User:     "postgres",  // Користувач
	Password: "postgres",  // Пароль
	Options:  map[string]string{"sslmode": "disable"},
}
