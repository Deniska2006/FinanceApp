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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

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
func InsertData(sess db.Session, data string) bool {
	var lastID sql.NullInt64
	row, _ := sess.SQL().QueryRow("SELECT MAX(id) FROM costs")
	err := row.Scan(&lastID)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal("Помилка виконання SQL:", err)
		return false
	}

	id := 1
	if lastID.Valid {
		id = int(lastID.Int64) + 1
	}

	name := Name(data)
	price := Price(data)
	if price == -1 {
		return false
	}

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
		return false
	}
	return true
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

	// InsertData(sess, "data")

	a := app.New()
	w := a.NewWindow("Вікно")
	w.Resize(fyne.NewSize(400, 320))

	label1 := widget.NewLabel("Вітаємо у фінаносвому трекері")
	entry1 := widget.NewEntry()
	label2 := widget.NewLabel("")

	data := ""
	isOk := false
	btn1 := widget.NewButton("Записати", func() {
		data = entry1.Text
		isOk = InsertData(sess, data)
		if isOk {
			label2.SetText("Дані записані")
		} else {
			label2.SetText("Помилка запису даних")
		}
		entry1.SetText("")

	})

	btn2 := widget.NewButton("Quit", func() {
		a.Quit()
	})

	w.SetContent(container.NewVBox(
		label1,

		entry1,
		btn1,
		btn2,
		label2,
	))
	w.ShowAndRun()

}
