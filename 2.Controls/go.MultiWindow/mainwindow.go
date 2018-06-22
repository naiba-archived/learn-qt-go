package main

import (
	"runtime"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//MultiWindow code gen ui
type MultiWindow struct {
	widgets.QMainWindow

	btn *widgets.QPushButton
}

func initMultiWindow() *MultiWindow {
	var this = NewMultiWindow(nil, 0)
	if runtime.GOOS == "darwin" {
		this.SetUnifiedTitleAndToolBarOnMac(true)
	}
	dbb := widgets.NewQDialog(this, 0)
	dbb.SetWindowTitle("Dialog Button Box")

	this.SetWindowTitle(core.QCoreApplication_ApplicationName())
	this.btn = widgets.NewQPushButton(this)
	this.btn.SetText("新窗口")
	this.btn.SetFixedWidth(100)
	this.btn.SetFixedHeight(50)

	this.btn.ConnectClicked(func(b bool) {
		this.SetWindowTitle(strconv.Itoa(dbb.Exec()))
	})

	return this
}
