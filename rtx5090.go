package main

import (
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Установщик RTX 5090")
	w.Resize(fyne.NewSize(400, 300))

	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Введите логин")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Введите пароль")

	progress := widget.NewProgressBar()
	progress.SetValue(0)

	statusLabel := widget.NewLabel("")

	startButton := widget.NewButton("Авторизироваться и установить", func() {
		login := loginEntry.Text
		password := passwordEntry.Text

		if login == "" {
			dialog.ShowError(
				errors.New("Неверный логин"),
				w,
			)
			return
		}

		if password != "qwerty123" {
			dialog.ShowError(
				errors.New("Неверный пароль"),
				w,
			)
			return
		}

		progress.SetValue(0)
		statusLabel.SetText("Установка...")

		go func() {
			for i := 1; i <= 20; i++ {
				time.Sleep(time.Second)
				progress.SetValue(float64(i) / 20.0)
			}
			statusLabel.SetText("Установка завершена! RTX 5090 теперь доступен на вашем компьютере!")
			dialog.ShowInformation("Успешно!", "Установка успешно завершена.", w)
		}()
	})

	form := container.NewVBox(
		widget.NewLabel("Логин:"),
		loginEntry,
		widget.NewLabel("Пароль:"),
		passwordEntry,
		startButton,
		progress,
		statusLabel,
	)

	w.SetContent(form)
	w.ShowAndRun()
}
