package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func makeBanner() fyne.CanvasObject {
    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.HomeIcon(), func() {}),
    )
    logo := canvas.NewImageFromResource(resourceLogoPng)
    logo.FillMode = canvas.ImageFillContain

    return container.NewStack(toolbar, logo)
}

func makeGui() fyne.CanvasObject {
    top := makeBanner()
    left := widget.NewLabel("Left")
    right := widget.NewLabel("Right")

    content := canvas.NewRectangle(color.Gray{Y: 0xee})

    dividers := [3]fyne.CanvasObject{
        widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
    }
    objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}
    return container.New(newChatLayout(top, left, right, content, dividers), objs...)
}

