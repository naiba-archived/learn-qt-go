package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var qApp *widgets.QApplication

func main() {
	qApp = widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("multi window")
	core.QCoreApplication_SetApplicationName("multi window")
	core.QCoreApplication_SetApplicationVersion("1.0.0")

	var mw = initMultiWindow()
	var availableGeometry = widgets.QApplication_Desktop().AvailableGeometry2(mw)
	mw.Resize2(600, 450)
	mw.Move2((availableGeometry.Width()-mw.Width())/2,
		(availableGeometry.Height()-mw.Height())/2)

	mw.Show()
	qApp.Exec()
}
