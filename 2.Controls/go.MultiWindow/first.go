package main

import (
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//MultiWindow code gen ui
type FirstWindow struct {
	widgets.QMainWindow

	btn *widgets.QPushButton
}

func initFirstWindow() *FirstWindow {
	var this = NewFirstWindow(nil, 0)
	var second = initSecondWindow(this, 0)

	if runtime.GOOS == "darwin" {
		this.SetUnifiedTitleAndToolBarOnMac(true)
	}
	this.SetWindowTitle(core.QCoreApplication_ApplicationName())
	this.SetFixedWidth(200)
	this.SetFixedHeight(100)

	this.btn = widgets.NewQPushButton(this)
	this.btn.SetText("打开新窗口")
	this.btn.SetFixedWidth(100)
	this.btn.SetFixedHeight(50)

	this.btn.ConnectClicked(func(b bool) {
		if widgets.QDialog__DialogCode(second.Exec()) == widgets.QDialog__Accepted {
			this.SetWindowTitle("Accept")
		} else {
			this.SetWindowTitle("Reject")
		}
	})

	return this
}
