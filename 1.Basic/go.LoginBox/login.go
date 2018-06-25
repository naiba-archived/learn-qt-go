package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//LoginWindow 登陆窗口
type LoginWindow struct {
	widgets.QDialog

	user *widgets.QLineEdit
	pass *widgets.QLineEdit
	btn  *widgets.QPushButton
}

func initLoginWindow() *LoginWindow {
	var this = NewLoginWindow(nil, core.Qt__WindowFullscreenButtonHint|core.Qt__WindowMaximizeButtonHint)
	this.SetMinimumSize2(150, 90)
	this.SetMaximumSize2(150, 90)

	this.user = widgets.NewQLineEdit(this)
	this.user.Move2(10, 10)
	this.pass = widgets.NewQLineEdit(this)
	this.pass.Move2(10, 40)
	this.btn = widgets.NewQPushButton(this)
	this.btn.Move2(40, 60)

	this.user.SetPlaceholderText("用户名")
	this.pass.SetEchoMode(widgets.QLineEdit__Password)
	this.pass.SetPlaceholderText("密码")
	this.btn.SetText("登陆")
	this.btn.ConnectClicked(func(c bool) {
		if this.user.Text() == "naiba" && this.pass.Text() == "123456" {
			this.Accept()
		} else {
			mb := widgets.NewQMessageBox(this)
			mb.Warning(this, "温馨提示", "账号或密码不正确", widgets.QMessageBox__Default, widgets.QMessageBox__Yes)
		}
	})

	return this
}
