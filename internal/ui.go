package internal

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"

	"github.com/upper/db/v4"
)

func AppMain(sess db.Session) {
	a := app.New()
	w := a.NewWindow("finTracker")
	w.Resize(fyne.NewSize(400, 300))

	showLogin(w, sess, a)

	w.ShowAndRun()
}

func showLogin(w fyne.Window, sess db.Session, a fyne.App) {
	mail := widget.NewEntry()
	password := widget.NewPasswordEntry()
	label := widget.NewLabel("Вхід у систему")

	loginButton := widget.NewButton("Увійти", func() {
		stratus, id := LogCheck(sess, mail.Text, password.Text)
		if stratus {
			Ui(w, sess, a, id)
		} else {
			label.SetText("Невірний логін або пароль")
			mail.SetText("")
			password.SetText("")
		}
	})

	registerButton := widget.NewButton("Реєстрація", func() {

		showRegister(w, sess, a)
	})

	quitButton := widget.NewButton("Вийти", func() {
		a.Quit()
	})

	w.SetContent(container.NewVBox(
		label,
		mail,
		password,
		loginButton,
		registerButton,
		quitButton,
	))
}

func showRegister(w fyne.Window, sess db.Session, a fyne.App) {
	newUser := widget.NewEntry()
	newPass := widget.NewPasswordEntry()
	confirm := widget.NewPasswordEntry()
	label := widget.NewLabel("Реєстрація")

	register := widget.NewButton("Зареєструватися", func() {
		if newPass.Text == confirm.Text {
			status, text := Register(sess, newUser.Text, newPass.Text)
			if !status {
				label.SetText(text)
			} else {
				showLogin(w, sess, a)
			}
		} else {
			w.SetContent(container.NewVBox(
				widget.NewLabel("Паролі не збігаються"),
				newUser, newPass, confirm,
				widget.NewButton("Зареєструватися", nil),
			))
		}
	})

	back := widget.NewButton("Назад", func() {
		showLogin(w, sess, a)
	})
	quitButton := widget.NewButton("Вийти", func() {
		a.Quit()
	})

	w.SetContent(container.NewVBox(
		label,
		newUser,
		newPass,
		confirm,
		register,
		back,
		quitButton,
	))
}

func Ui(w fyne.Window, sess db.Session, a fyne.App, uid int64) {

	labelС := widget.NewLabel("Впишіть назву витрати")
	labelM := widget.NewLabel("Скільки ви витратили")
	label1 := widget.NewLabel("Вітаємо у фінаносвому трекері")
	entryAdd := widget.NewEntry()
	entryAdd.PlaceHolder = "Впишіть категорію"
	labelData := widget.NewLabel("")
	label1.Alignment = fyne.TextAlignCenter
	entryM := widget.NewEntry()
	label2 := widget.NewLabel("")

	dropdown := widget.NewSelect(GetCategories(sess, uid), func(selected string) {})
	dropdown.PlaceHolder = "Оберіть категорію"

	isOk := false
	btn1 := widget.NewButton("Записати", func() {
		isOk = InsertData(sess, dropdown.Selected, entryM.Text, uid)
		if isOk {
			label2.SetText("Дані записані")
		} else {
			label2.SetText("Помилка запису даних")
		}
		entryM.SetText("")
		dropdown.ClearSelected()
	})

	btn2 := widget.NewButton("Вийти", func() {
		a.Quit()
	})
	var btn4, btn3, btn5, btnAdd *widget.Button

	btn4 = widget.NewButton("Сховати", func() {
		label2.SetText("")
		btn4.Hide()
		labelData.Hide()
		w.Content().Refresh()
		time.Sleep(50 * time.Millisecond)
		w.Resize(fyne.NewSize(400, 300))

	})
	btn3 = widget.NewButton("Показати дані", func() {
		label2.SetText("")
		btn4.Show()
		labelData.Show()
		labelData.SetText(GetData(sess, uid))
	})
	btn5 = widget.NewButton("Додати категорію", func() {
		label1.Hide()
		labelС.Hide()
		dropdown.Hide()
		labelM.Hide()
		entryM.Hide()
		btn1.Hide()
		btn3.Hide()
		btn5.Hide()
		label2.Hide()
		labelData.Hide()
		btn4.Hide()
		entryAdd.Show()
		btnAdd.Show()

	})
	btnAdd = widget.NewButton("Додати", func() {
		if entryAdd.Text != "" {
			AddCategory(sess, entryAdd.Text, uid)
		}
		labelС.Show()
		dropdown.Show()
		labelM.Show()
		entryM.Show()
		btn1.Show()
		btn3.Show()
		btn5.Show()
		label1.Show()
		label2.Show()
		entryAdd.Hide()
		btnAdd.Hide()
		entryAdd.SetText("")
		entryAdd.PlaceHolder = "Впишіть категорію"
		newCategories := GetCategories(sess, uid)
		dropdown.Options = newCategories
		dropdown.Refresh()

	})
	entryAdd.Hide()
	labelData.Hide()
	btnAdd.Hide()
	btn4.Hide()
	w.SetContent(container.NewVBox(
		label1,
		labelС,
		dropdown,
		labelM,
		entryM,
		btn1,
		btn3,
		btn5,
		entryAdd,
		btnAdd,
		btn2,
		label2,
		labelData,
		btn4,
	))

}
