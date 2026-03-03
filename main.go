package main

import (
	"fmt"

	"github.com/qqxaGo/passChecker/checker"
)

func main() {
	var password string

	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	level, missing := checker.Analyze(password)

	fmt.Println("Уровень:", level)

	if len(missing) > 0 {
		fmt.Println("Что улучшить: ")
		for _, m := range missing {
			fmt.Println("-", m)
		}
	}
}
