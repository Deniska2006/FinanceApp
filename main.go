package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       = 100
	pointsPerQuestion = 20
)

func main() {
	fmt.Println("Вітаємо у грі MathCore!")

	for i := 5; i > 0; i-- {
		fmt.Printf("До початку: %v c\n", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У Вас %v очок!\n", myPoints)
			} else {
				fmt.Println("НЕ ПРАВИЛЬНО!")
			}
		}
	}
}
