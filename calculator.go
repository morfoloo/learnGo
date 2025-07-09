package main

import (
	"fmt"
	"math"
)

func main() {
	var op string
	for {
		fmt.Println("\nВведите операцию: +, -, *, /, sqrt(Кв. корень), sin(Синус), cos(Косинус), tan(Тангенс), pow(Возведение в степень). Используйте exit или quit для выхода.")
		fmt.Print("> ")
		fmt.Scan(&op)

		if op == "exit" || op == "quit" {
			fmt.Println("Выход...")
			break
		}
		switch op {
		case "+", "-", "*", "/":
			var a, b float64
			fmt.Print("Введите два числа: ")
			fmt.Scan(&a, &b)

			switch op {
			case "+":
				fmt.Printf("%v + %v = %.4f\n", a, b, a+b)
			case "-":
				fmt.Printf("%v - %v = %.4f\n", a, b, a-b)
			case "*":
				fmt.Printf("%v * %v = %.4f\n", a, b, a*b)
			case "/":
				if b != 0 {
					fmt.Printf("%v + %v = %.4f\n", a, b, a/b)
				} else {
					fmt.Println("Не поделю на ноль.")
				}
			}
		case "sqrt":
			var a float64
			fmt.Print("Введите число: ")
			fmt.Scan(&a)
			if a > 0 {
				fmt.Printf("Квадратный корень из %v равен %.4f\n", a, math.Sqrt(a))
			} else {
				fmt.Println("Введите число больше нуля!")
			}

		case "sin":
			var a float64
			fmt.Print("Введите угол в радианах: ")
			fmt.Scan(&a)
			fmt.Printf("Синус угла %v равен %.4f\n", a, math.Sin(a))
		case "cos":
			var a float64
			fmt.Print("Введите угол в радианах: ")
			fmt.Scan(&a)
			fmt.Printf("Косинус угла %v равен %.4f\n", a, math.Cos(a))
		case "tan":
			var a float64
			fmt.Print("Введите угол в радианах: ")
			fmt.Scan(&a)
			fmt.Printf("Тангенс угла %v равен %.4f\n", a, math.Tan(a))
		case "pow":
			var a, b float64
			fmt.Print("Введите число и степень: ")
			fmt.Scan(&a, &b)
			fmt.Printf("%v в степени %v = %.4f\n", a, b, math.Pow(a, b))

		default:
			fmt.Println("Не распознал операцию.")
		}
	}
}
