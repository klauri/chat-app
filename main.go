package main

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(newChatTheme())
    w := a.NewWindow("Login")
    w.Resize(fyne.NewSize(1024, 768))
    
    ui := &gui{win: w}
    w.SetContent(ui.makeGui())
    w.SetMainMenu(ui.makeMenu())
    
    dirPath := "C:\\Users\\kirkl"
    dirPath, err := filepath.Abs(dirPath)
    if err != nil {
        fmt.Println("Error resolving project path", err)
        return
    }

    dirURI := storage.NewFileURI(dirPath)
    dir, err := storage.ListerForURI(dirURI)
    if err != nil {
        fmt.Println("Error opening project", nil)
        return
    }

    ui.openProject(dir)
    ui.showCreate(w)
	w.ShowAndRun()
}

func (g *gui) makeMenu() *fyne.MainMenu {
    file := fyne.NewMenu("File",
        fyne.NewMenuItem("Open Project", g.openProjectDialog),
    )

    return fyne.NewMainMenu(file)
}
