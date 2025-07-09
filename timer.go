package main

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	for {
		var ch string
		fmt.Println(aurora.Bold(aurora.Cyan("Добро пожаловать в таймер! Выбери:\n1 - Таймер\n2 - Секундомер\nq - Выйти")))
		fmt.Scan(&ch)

		switch ch {
		case "1":
			var sec int
			fmt.Print(aurora.Bold(aurora.Magenta("Ты попал в таймер! Введите время в секундах: ")))
			fmt.Scan(&sec)

			for i := sec; i > 0; i-- {
				fmt.Println(aurora.Bold(aurora.Green(i)))
				time.Sleep(1 * time.Second)
				clear()
			}
			fmt.Println(aurora.Bold(aurora.Blue("Время вышло!")))
		case "2":
			fmt.Println(aurora.Bold(aurora.Yellow("Ты попал в секундомер! Нажми Enter чтобы начать!")))
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			start := time.Now()
			stop := make(chan bool)

			go func() {
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				stop <- true
			}()

			fmt.Println(aurora.Bold(aurora.Green("▶ Секундомер запущен! Нажми Enter, чтобы остановить.")))

		loop:
			for {
				select {
				case <-stop:
					elapsed := time.Since(start)
					fmt.Printf("\n⏹ Секундомер остановлен! Прошло времени: %02d:%02d\n",
						aurora.Bold(aurora.Green(int(elapsed.Minutes()))), aurora.Bold(aurora.Green(int(elapsed.Seconds())%60)))
					break loop
				default:
					elapsed := time.Since(start)
					fmt.Printf("\rПрошло: %02d:%02d", aurora.Bold(aurora.Blue(int(elapsed.Minutes()))), aurora.Bold(aurora.Blue(int(elapsed.Seconds())%60)))
					time.Sleep(1 * time.Second)
				}
			}
		case "q":
			fmt.Println(aurora.Bold(aurora.Cyan("Пока!")))
			return
		default:
			fmt.Println(aurora.Bold(aurora.Red("Неверный выбор!")))
		}
	}
}
