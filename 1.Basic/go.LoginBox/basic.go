package main

import (
	"github.com/therecipe/qt/widgets"
)

//BasicWindow 主窗口
type BasicWindow struct {
	widgets.QMainWindow

	hello *widgets.QLabel
}

func initBasicWindow() *BasicWindow {
	var this = NewBasicWindow(nil, 0)
	this.SetWindowTitle("主窗口")
	this.SetFixedSize2(300, 50)
	this.hello = widgets.NewQLabel(this, 0)

	this.hello.SetText("hello")

	return nil
}
