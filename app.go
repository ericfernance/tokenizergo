package main

import (
	_ "embed"
	"fmt"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	_ "github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/gtkutil"
	"os"
)

//go:embed tokenizer.ui
var windowUi string

type App struct {
	Application *gtk.Application
}

func NewApp() *App {
	app := &App{
		Application: gtk.NewApplication("com.thisisericrobert.tokenizer", gio.ApplicationFlagsNone),
	}
	app.Application.ConnectActivate(app.activate)
	app.Application.Run(os.Args)
	return app
}

func (a *App) activate() {
	fmt.Println("Hello, gotk4!")
	builder := gtk.NewBuilderFromString(windowUi, len(windowUi))
	window := builder.GetObject("MainWindow").Cast().(*gtk.ApplicationWindow)
	menuButton := builder.GetObject("topMenu").Cast().(*gtk.MenuButton)
	quitAction := gio.NewSimpleAction("quit", nil)
	saveAction := gio.NewSimpleAction("save", nil)
	a.Application.AddAction(quitAction)
	a.Application.AddAction(saveAction)
	quitAction.ConnectActivate(func(param *glib.Variant) {
		fmt.Println("Action activated!")
		a.Application.Quit()
	})
	saveAction.ConnectActivate(func(param *glib.Variant) {
		fmt.Println("Save action activated!")
	})
	a.Application.SetAccelsForAction("app.quit", []string{"<Ctrl>Q"})
	a.Application.SetAccelsForAction("app.save", []string{"<Ctrl>S"})
	menuButton.SetMenuModel(gtkutil.MenuPair([][2]string{{"Quit", "app.quit"}}))
	window.SetApplication(a.Application)
	window.Show()
}

func (a *App) encode() {
	fmt.Println("Encode!")
}

func (a *App) decode() {
	fmt.Println("Decode!")
}
