package main

import (
	"chat-app/internal/dialogs"
	"errors"
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
    logo := canvas.NewImageFromResource(resourceAssetsLogoPng)
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
        wizard.Push("Step 2", g.makeCreateDetail(wizard))
    })
    create.Importance = widget.HighImportance

    buttons := container.NewGridWithColumns(2, open, create)
    home := container.NewVBox(intro, buttons)

    wizard = dialogs.NewWizard("Create Project", home)
    wizard.Show(w)
    wizard.Resize(home.MinSize().AddWidthHeight(40, 80))  //fyne.NewSize(360,200))
}

func (g *gui) makeCreateDetail(wizard *dialogs.Wizard) fyne.CanvasObject {
    homeDir, _ := os.UserHomeDir()
    parent := storage.NewFileURI(homeDir)
    chosen, _ := storage.ListerForURI(parent)

    name := widget.NewEntry()
    name.Validator = func(in string) error {
        if in == "" {
            return errors.New("Project name is required")
        }

        return nil
    }
    var dir *widget.Button
    dir = widget.NewButton(chosen.Name(), func() {
        d := dialog.NewFolderOpen(func(l fyne.ListableURI, err error) {
            if err != nil || l == nil {
                return
            }

            chosen = l

            dir.SetText(l.Name())
        }, g.win)

        d.SetLocation(chosen)
        d.Show()
    })

    form := widget.NewForm(
        widget.NewFormItem("Name", name),
        widget.NewFormItem("Parent Directory", dir),
    )
<<<<<<< HEAD

    form.OnSubmit = func() {
        if name.Text == "" {
            return
        }

        project, err := createProject(name.Text, chosen)
        if err != nil {
            dialog.ShowError(err, g.win)
            return
        }
        wizard.Hide()
        g.openProject(project)
    }

=======
>>>>>>> 322ff84ffd84ea8320dd2440e55b02073482fe66
    return form
}

