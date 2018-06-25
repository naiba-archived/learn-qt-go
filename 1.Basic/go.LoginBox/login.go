package main

import (
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
	var this = NewLoginWindow(nil, 0)
	this.SetMinimumSize2(147, 93)
	this.SetMaximumSize2(147, 93)
	this.SetWindowTitle("用户登陆")

	this.user = widgets.NewQLineEdit(this)
	this.user.Move2(10, 10)
	this.pass = widgets.NewQLineEdit(this)
	this.pass.Move2(10, 40)
	this.btn = widgets.NewQPushButton(this)
	this.btn.Move2(40, 63)

	this.user.SetPlaceholderText("用户名")
	this.pass.SetEchoMode(widgets.QLineEdit__Password)
	this.pass.SetPlaceholderText("密码")
	this.btn.SetText("登陆")
	this.btn.ConnectClicked(func(c bool) {
		if this.user.Text() == "naiba" && this.pass.Text() == "123456" {
			this.Accept()
		} else {
			mb := widgets.NewQMessageBox(this)
			mb.Warning(this, "温馨提示", "账号或密码不正确:"+this.user.Text()+","+this.pass.Text(), widgets.QMessageBox__Default, widgets.QMessageBox__Yes)
		}
	})

	return this
}
