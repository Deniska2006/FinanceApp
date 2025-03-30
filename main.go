package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Id() string {
	file, err := os.Open("finance.txt")
	if err != nil {
		log.Print("Помилка відкриття файлу:", err)
		return "1"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastLine string

	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if lastLine == "" {
		return "1"
	}

	parts := strings.Split(lastLine, " ")
	if len(parts) == 0 {
		return "1"
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("Не коректний id у файлі: %v", err)
		return "error"
	}

	id++
	return strconv.Itoa(id)
}

func InsertData(data string) {

	file, err := os.OpenFile("finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Print("Помилка відкриття файлу:", err)
		return
	}

	defer file.Close()

	id := Id()

	date := time.Now()

	data = strings.ReplaceAll(data, " ", " \t")

	_, err = file.WriteString(id + " " + data + "\t" + date.Format("2006-01-02 15:04:05") + "\n")
	if err != nil {
		log.Print("Помилка запису у файл:", err)
		return
	}

}

func main() {
	var action uint8 = 10

	fmt.Println("0 - вихід, 1 - записати фін. опр.,2 - delete")

	scanner := bufio.NewScanner(os.Stdin)
	for action != 0 {
		fmt.Scan(&action)
		switch action {

		case 1:
			fmt.Scanln()
			scanner.Scan()
			data := scanner.Text()

			InsertData(data)
		case 2:
			os.Truncate("finance.txt", 0)

		}

	}

}
