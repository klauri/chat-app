package main

import (
	"chat-app/internal/dialogs"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
    win fyne.Window
    directory *widget.Label
}

func makeBanner() fyne.CanvasObject {
    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.HomeIcon(), func() {}),
    )
    logo := canvas.NewImageFromResource(resourceLogoPng)
    logo.FillMode = canvas.ImageFillContain

    return container.NewStack(toolbar, logo)
}

func (g *gui) makeGui() fyne.CanvasObject {
    top := makeBanner()
    left := widget.NewLabel("Left")
    right := widget.NewLabel("Right")
    
    g.directory = widget.NewLabel("Welcome - Please select a directory . . . ")
    content := container.NewStack(canvas.NewRectangle(color.Gray{Y: 0xee}), g.directory)

    dividers := [3]fyne.CanvasObject{
        widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
    }
    objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}
    return container.New(newChatLayout(top, left, right, content, dividers), objs...)
}

func (g *gui) openProjectDialog() {
    dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
        if err != nil {
            dialog.ShowError(err, g.win)
        }
        if dir == nil {
            return
        }
        
        g.openProject(dir)
    }, g.win) 
}

func (g *gui) openProject(dir fyne.ListableURI) {
        name := dir.Name()

        g.win.SetTitle("Fyne App: " + name)
        g.directory.SetText(name)

}

func (g *gui) showCreate(w fyne.Window) {
    var wizard *dialogs.Wizard
    intro := widget.NewLabel(`Here you can create a new project!

Or open an existing one that you created earlier.`)
    
    open := widget.NewButton("Open Project", func() {
        wizard.Hide()
        g.openProjectDialog()
    })
    create := widget.NewButton("Create Project", func() {
        wizard.Push("Step 2", g.makeCreateDetail())
    })
    create.Importance = widget.HighImportance

    buttons := container.NewGridWithColumns(2, open, create)
    home := container.NewVBox(intro, buttons)

    wizard = dialogs.NewWizard("Create Project", home)
    wizard.Show(w)
    wizard.Resize(home.MinSize().AddWidthHeight(40, 80))  //fyne.NewSize(360,200))
}

func (g *gui) makeCreateDetail() fyne.CanvasObject {
    homeDir, _ := os.UserHomeDir()
    parent := storage.NewFileURI(homeDir)
    chosen, _ := storage.ListerForURI(parent)

    name := widget.NewEntry()
    dir := widget.NewButton(chosen.Name(), func() {
        // TODO open dialog
    })

    form := widget.NewForm(
        widget.NewFormItem("Name", name),
        widget.NewFormItem("Parent Directory", dir),
    )
}

