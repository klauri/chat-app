package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(newChatTheme())
    w := a.NewWindow("Login")
    w.Resize(fyne.NewSize(1024, 768))
    
    w.SetContent(makeGui())
	w.ShowAndRun()
}
