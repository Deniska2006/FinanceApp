package dbm

import (
	"fmt"
	"io/ioutil"

	"github.com/upper/db/v4"
)

func ExecuteSQLFile(sess db.Session, filePath string) error {
	// Зчитуємо файл
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("не вдалося прочитати файл %s: %v", filePath, err)
	}

	// Виконуємо SQL-команди
	_, err = sess.SQL().Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("помилка виконання SQL-команд: %v", err)
	}

	return nil
}
