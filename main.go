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
    
    ui := &gui{win: w}
    w.SetContent(ui.makeGui())
    w.SetMainMenu(ui.makeMenu())

    ui.openProject()
	w.ShowAndRun()
}

func (g *gui) makeMenu() *fyne.MainMenu {
    file := fyne.NewMenu("File",
        fyne.NewMenuItem("Open Project", g.openProject),
    )

    return fyne.NewMainMenu(file)
}
