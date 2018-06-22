package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//SecondWindow code gen ui
type SecondWindow struct {
	widgets.QDialog

	btn *widgets.QPushButton
}

func initSecondWindow(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SecondWindow {
	var this = NewSecondWindow(parent, 0)
	this.SetWindowTitle("第二窗口")
	this.SetFixedWidth(200)
	this.SetFixedHeight(100)

	this.btn = widgets.NewQPushButton(this)
	this.btn.SetText("返回主窗口")
	this.btn.SetStyleSheet("color:black")
	this.btn.SetFixedWidth(100)
	this.btn.SetFixedHeight(50)

	this.btn.ConnectClicked(func(b bool) {
		this.Accept()
	})

	return this
}
