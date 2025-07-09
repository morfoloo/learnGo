package main

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"os"
	"os/exec"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearConsole()
		var ch string
		fmt.Println(aurora.Bold(aurora.Cyan("Добро пожаловать в конвертер температур! Выбери:\n1 - Из Цельсия в Фаренгейт\n2 - Из Фаренгейта в Цельсий\nq - Выйти")))
		fmt.Scan(&ch)
		switch ch {
		case "1":
			var celsium float64
			fmt.Print(aurora.Bold(aurora.Blue("Введи температуру в градусах Цельсия: ")))
			fmt.Scan(&celsium)
			farenheit := celsium*1.8 + 32
			fmt.Printf("%v градусов Цельсия = %v градусов Фаренгейта\n", aurora.Bold(aurora.Blue(celsium)), aurora.Bold(aurora.Green(farenheit)))
		case "2":
			var farenheit float64
			fmt.Print(aurora.Bold(aurora.Green("Введи температуру в градусах Фаренгейта: ")))
			fmt.Scan(&farenheit)
			celsium := (farenheit - 32) * 5 / 9
			fmt.Printf("%v градусов Фаренгейта = %v градусов Цельсия\n", aurora.Bold(aurora.Green(farenheit)), aurora.Bold(aurora.Green(celsium)))
		case "q":
			fmt.Println(aurora.Bold(aurora.Green("Пока!!")))
			return
		default:
			fmt.Println(aurora.Bold(aurora.Red("Неверный выбор! Попробуй снова!")))
		}
	}
}
