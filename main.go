package main

import (
	"database/sql"
	"fmt"
	"log"
	"prj-test/config"
	"prj-test/dbm"
	"prj-test/domain"
	"strconv"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

// func Id() int {
// 	file, err := os.Open("finance.txt")
// 	if err != nil {
// 		log.Print("Помилка відкриття файлу:", err)
// 		return 0
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var lastLine string

// 	for scanner.Scan() {
// 		lastLine = scanner.Text()
// 	}

// 	if lastLine == "" {
// 		return 1
// 	}

// 	idS := ""
// 	for _, v := range lastLine {
// 		if string(v) == " " || string(v) == "\t" {
// 			break
// 		}
// 		idS += string(v)
// 	}
// 	id, err := strconv.Atoi(idS)
// 	if err != nil {
// 		log.Printf("Не коректний id у файлі: %v", err)
// 		return -1
// 	}

// 	id++
// 	return id
// }

func Price(data string) int {
	price := -1
	for i, v := range data {
		if string(v) == " " {
			var err error
			price, err = strconv.Atoi(data[i+1:])
			if err != nil {
				fmt.Println("Помилка конвертації")
				return -1
			}
			return price

		}
	}
	return price
}

func Name(data string) string {
	name := ""
	for _, v := range data {
		if string(v) == " " {
			break
		}
		name += string(v)

	}
	return name
}
func InsertData(sess db.Session, data string) {

	// file, err := os.OpenFile("finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	// if err != nil {
	// 	log.Print("Помилка відкриття файлу:", err)
	// 	return
	// }

	// defer file.Close()

	var lastID sql.NullInt64
	row, _ := sess.SQL().QueryRow("SELECT MAX(id) FROM costs")
	err := row.Scan(&lastID)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal("Помилка виконання SQL:", err)
	}

	id := 1
	if lastID.Valid {
		id = int(lastID.Int64) + 1
	}

	name := Name(data)
	price := Price(data)

	date := time.Now()

	product := domain.Cost{
		ID:          id,
		Name:        name,
		Price:       price,
		CreatedTime: date.Format("2006-01-02 15:04:05"),
	}
	collection := sess.Collection("costs")
	_, err = collection.Insert(&product)
	if err != nil {
		log.Fatal("Error inserting data: ", err)
	}

	// _, err = file.WriteString(strconv.Itoa(id) + "\t" + name + "\t" + strconv.Itoa(price) + "\t" + date.Format("2006-01-02 15:04:05") + "\n")
	// if err != nil {
	// 	log.Print("Помилка запису у файл:", err)
	// 	return
	// }

}

func main() {

	sess, err := postgresql.Open(config.Settings)
	if err != nil {
		log.Fatal("Не вдалося підключитися до бази:", err)
	}
	defer sess.Close()

	fmt.Println("✅ Успішно підключено до бази!")

	err = dbm.ExecuteSQLFile(sess, "dbm\\schema.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу:", err)
	} else {
		fmt.Println("Tables being create")
	}

	InsertData(sess, "cock 300")

	// var action uint8 = 10

	// fmt.Println("0 - вихід, 1 - записати фін. опр.,2 - delete")

	// scanner := bufio.NewScanner(os.Stdin)
	// for action != 0 {
	// 	fmt.Scan(&action)
	// 	switch action {

	// 	case 1:
	// 		fmt.Scanln()
	// 		scanner.Scan()
	// 		data := scanner.Text()

	// 		InsertData(sess, data)
	// 	case 2:
	// 		os.Truncate("finance.txt", 0)

	// 	}

	// }

}
