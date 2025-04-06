package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
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

func GetData(sess db.Session) string {
	var all float64
	var costs []domain.Data
	err := sess.Collection("costs").Find().All(&costs)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]float64)

	for _, cost := range costs {
		all += cost.Price
		m[cost.Name] += cost.Price

	}
	result := "Ви витратили:\n"
	for i, v := range m {
		result += fmt.Sprintf("На %s - %.2fгрн(%.2f%%)\n ", i, v, (100*v)/all)

	}
	result += fmt.Sprintf("Всього - %fгрн\n", all)
	return result
}

func GetCategories(sess db.Session) []string {
	var categories []map[string]interface{}
	err := sess.Collection("categories").Find().All(&categories)
	if err != nil {
		log.Fatal(err)
	}

	var categoryNames []string
	for _, category := range categories {

		if name, ok := category["name"].(string); ok {
			categoryNames = append(categoryNames, name)
		}
	}
	return categoryNames

}

func InsertData(sess db.Session, name string, price string) bool {
	if name == "" {
		return false
	}

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

	price1, err := strconv.ParseFloat(price, 32)
	price1 = math.Round(price1*100) / 100

	if err != nil {
		return false
	}

	date := time.Now()

	product := domain.Cost{
		ID:          id,
		Name:        name,
		Price:       price1,
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

	err = dbm.ExecuteSQLFile(sess, "dbm\\categories.sql")
	if err != nil {
		log.Fatal("Помилка при виконанні SQL-файлу:", err)
	} else {
		fmt.Println("Tables being create2")
	}

	a := app.New()
	w := a.NewWindow("FinanceTraker")
	w.Resize(fyne.NewSize(400, 400))

	labelС := widget.NewLabel("Впишіть назву витрати")
	labelM := widget.NewLabel("Скільки ви витратили")
	label1 := widget.NewLabel("Вітаємо у фінаносвому трекері")
	labelData := widget.NewLabel("")
	labelData.Hide()
	label1.Alignment = fyne.TextAlignCenter
	entryM := widget.NewEntry()

	label2 := widget.NewLabel("")

	options := GetCategories(sess)

	dropdown := widget.NewSelect(options, func(selected string) {})
	dropdown.PlaceHolder = "Оберіть категорію"

	isHide := true
	isOk := false
	btn1 := widget.NewButton("Записати", func() {
		isOk = InsertData(sess, dropdown.Selected, entryM.Text)
		if isOk {
			label2.SetText("Дані записані")
		} else {
			label2.SetText("Помилка запису даних")
		}
		entryM.SetText("")
		dropdown.ClearSelected()
	})

	btn2 := widget.NewButton("Вийти", func() {
		a.Quit()
	})
	var btn4, btn3 *widget.Button
	updateButtons := func() {
		if isHide {
			w.Resize(fyne.NewSize(400, 400))
			btn4.Hide()
			labelData.Hide()

		} else {
			w.Resize(fyne.NewSize(400, 500))
			btn4.Show()
			labelData.Show()

		}
	}
	btn4 = widget.NewButton("Сховати", func() {
		label2.SetText("")
		isHide = true
		updateButtons()
	})
	btn3 = widget.NewButton("Показати дані", func() {
		label2.SetText("")
		isHide = false
		updateButtons()
		labelData.SetText(GetData(sess))
	})
	updateButtons()

	w.SetContent(container.NewVBox(
		label1,
		labelС,
		dropdown,
		labelM,
		entryM,
		btn1,
		btn3,
		btn2,
		label2,
		labelData,
		btn4,
	))
	w.ShowAndRun()

}
