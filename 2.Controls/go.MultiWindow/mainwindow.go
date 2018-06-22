package main

import (
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//MultiWindow code gen ui
type MultiWindow struct {
	widgets.QMainWindow

	label *widgets.QLabel
}

func initMultiWindow() *MultiWindow {
	var this = NewMultiWindow(nil, 0)
	if runtime.GOOS == "darwin" {
		this.SetUnifiedTitleAndToolBarOnMac(true)
	}
	this.SetWindowTitle(core.QCoreApplication_ApplicationName())
	this.label = widgets.NewQLabel2("hello world!", this, 0)
	this.label.SetStyleSheet("background-color: red")
	return this
}
