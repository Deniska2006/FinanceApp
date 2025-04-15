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

	err = dbm.ExecuteSQLFile(sess, "../dbm/users.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу  /dbm/users.sql : ", err)
	}

	err = dbm.ExecuteSQLFile(sess, "../dbm/schema.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу  /dbm/schema.sql  :", err)
	}

	err = dbm.ExecuteSQLFile(sess, "../dbm/categories.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу  /dbm/categories.sql :  ", err)
	}

	internal.AppMain(sess)
}
