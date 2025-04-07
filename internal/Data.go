package internal

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"something/domain"
	"strconv"
	"time"

	"github.com/upper/db/v4"
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
	result += fmt.Sprintf("Всього - %.2fгрн\n", all)
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

func AddCategory(sess db.Session, c string) {
	category := domain.Category{Name: c}
	err := sess.Collection("categories").InsertReturning(&category)
	if err != nil {
		log.Fatal("Insert error:", err)
	}
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
