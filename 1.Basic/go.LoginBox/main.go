package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func init() {
	core.QCoreApplication_SetOrganizationName("奶爸")
	core.QCoreApplication_SetApplicationName("登陆窗口")
	core.QCoreApplication_SetApplicationVersion("1.0.0")
}

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	var mw = initBasicWindow()
	var login = initLoginWindow()

	if widgets.QDialog__DialogCode(login.Exec()) == widgets.QDialog__Accepted {
		mw.Show()
		qApp.Exec()
	} else {
		qApp.Quit()
	}
}
