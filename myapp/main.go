package main

import (
	"log"
	"something/config"
	"something/dbm"
	"something/internal"

	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {

	sess, err := postgresql.Open(config.Settings)
	if err != nil {
		log.Fatal("Не вдалося підключитися до бази:", err)
	}
	defer sess.Close()

	err = dbm.ExecuteSQLFile(sess, "../dbm/schema.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу:", err)
	}

	err = dbm.ExecuteSQLFile(sess, "../dbm/categories.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу:", err)
	}

	internal.Ui(sess)
}
