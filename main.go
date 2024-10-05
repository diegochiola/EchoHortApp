package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App        fyne.App    //referencia a la app o canva
	InfoLog    *log.Logger // Log de acciones ,puntero por eso el *
	ErrorLog   *log.Logger // Log de errores
	MainWindow fyne.Window // ventana principal HOME

}

var myApp Config // variable global que contiene el struct config

func main() {
	//creacion de la App
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp") //identificador unico
	myApp.App = fyneApp
	//crear logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  //plantilla para registros de logs
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) //plantilla para errores
	// Conexión a base de datos

	//creacion del repo de base de datos

	//Definir tamaño y caracteristicas del home
	myApp.MainWindow = fyneApp.NewWindow("ECO HORT APP")
	myApp.MainWindow.Resize(fyne.NewSize(800, 500)) //tamano de la ventana
	myApp.MainWindow.SetFixedSize(true)             //para que no se pueda redimensionar
	myApp.MainWindow.SetMaster()                    //para que sea la ventana principal
	//mostrar y ejecutar la app
	myApp.MainWindow.ShowAndRun()
	/*
		myApp := app.New()
		window := myApp.NewWindow("ECO HORT APP")
		window.SetContent(widget.NewLabel("Welcome to myApp!")) //ventana quedara del tamaño del texto
		window.ShowAndRun()                                     //muestra y ejecuta la ventana
	*/

}
