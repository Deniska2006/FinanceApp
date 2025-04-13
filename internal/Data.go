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
	"golang.org/x/crypto/bcrypt"
)

func LogCheck(sess db.Session, login, password string) (bool, int64) {
	var user domain.Account

	// Знайти користувача по логіну
	err := sess.Collection("users").Find(db.Cond{"mail": login}).One(&user)
	if err != nil {
		log.Println("Користувача не знайдено:", err)
		return false, 0
	}

	// Порівняти хешований пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		log.Println("Невірний пароль")
		return false, 0
	}

	return true, user.Id
}

func Register(sess db.Session, mail, password string) (bool, string) {

	var existingUser domain.Account
	err := sess.Collection("users").Find(db.Cond{"mail": mail}).One(&existingUser)
	if err == nil {

		log.Println("Користувач з таким email вже існує.")
		return false, "Користувач з таким email вже існує."
	}

	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Помилка хешування пароля:", err)
		return false, "Помилка хешування пароля"
	}

	user := domain.Account{
		Mail:           mail,
		HashedPassword: string(h),
	}

	collection := sess.Collection("users")
	_, err = collection.Insert(&user)
	if err != nil {
		log.Println("Помилка вставки даних:", err)
		return false, "Помилка вставки даних"
	}

	return true, ""
}

func GetData(sess db.Session, uid int64) string {
	var all float64
	var costs []domain.Data
	err := sess.Collection("costs").Find(db.Cond{"uid": uid}).All(&costs)
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

func GetCategories(sess db.Session, uid int64) []string {
	var categories []domain.Category
	err := sess.Collection("categories").Find(db.Cond{"uid": uid}).All(&categories)
	if err != nil {
		log.Fatal(err)
	}

	var categoryNames []string
	for _, category := range categories {
		if category.Uid == uid {
			categoryNames = append(categoryNames, category.Name)
		}
	}
	return categoryNames

}

func AddCategory(sess db.Session, c string, uid int64) {
	category := domain.Category{Name: c, Uid: uid}
	err := sess.Collection("categories").InsertReturning(&category)
	if err != nil {
		log.Fatal("Insert error:", err)
	}
}

func InsertData(sess db.Session, name string, price string, uid int64) bool {
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
		Uid:         int(uid),
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
