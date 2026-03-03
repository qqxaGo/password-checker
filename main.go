package main

import (
	"fmt"

	"github.com/qqxaGo/passChecker/checker"
)

func main() {
	var password string

	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	result := checker.Analyze(password)

	fmt.Println("Результат:", result)
}
