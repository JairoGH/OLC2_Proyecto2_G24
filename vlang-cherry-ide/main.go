package main

import (
    "fmt"
    "main/ui"
    "fyne.io/fyne/v2/app"
)

func main() {
    // Limpiar consola usando códigos ANSI
    fmt.Print("\033[H\033[2J")
    
    // Crear aplicación
    myApp := app.New()
    myApp.SetIcon(nil) // Puedes agregar un icono más tarde

    // Crear y mostrar ventana principal
    mainWindow := ui.NewMainWindow(myApp)
    mainWindow.ShowAndRun()
}
