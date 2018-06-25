package main

import (
	"os"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//Hello code gen ui
type Hello struct {
	widgets.QMainWindow

	label *widgets.QLabel
}

var qApp *widgets.QApplication

func main() {
	qApp = widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("hello world")
	core.QCoreApplication_SetApplicationName("hello world")
	core.QCoreApplication_SetApplicationVersion("1.0.0")

	var mw = initHello()
	var availableGeometry = widgets.QApplication_Desktop().AvailableGeometry2(mw)
	mw.Resize2(600, 450)
	mw.Move2((availableGeometry.Width()-mw.Width())/2,
		(availableGeometry.Height()-mw.Height())/2)

	mw.Show()
	qApp.Exec()
}

func initHello() *Hello {
	var this = NewHello(nil, 0)
	if runtime.GOOS == "darwin" {
		this.SetUnifiedTitleAndToolBarOnMac(true)
	}
	this.SetWindowTitle(core.QCoreApplication_ApplicationName())
	this.label = widgets.NewQLabel2("hello world!", this, 0)
	this.label.SetStyleSheet("background-color: red")
	return this
}
