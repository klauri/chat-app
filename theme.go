//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type chatTheme struct {
	fyne.Theme
}

func newChatTheme() fyne.Theme {
    return &chatTheme{Theme: theme.DefaultTheme()}
}

func (t *chatTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
    return t.Theme.Color(name, theme.VariantLight)
}

func (t *chatTheme) Size(name fyne.ThemeSizeName) float32 {
    if name == theme.SizeNameText {
        return 12
    }

    return t.Theme.Size(name)
}

