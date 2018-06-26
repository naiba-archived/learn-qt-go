package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func init() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	core.QCoreApplication_SetOrganizationName("奶爸")
	core.QCoreApplication_SetApplicationName("图形视图框架")
	core.QCoreApplication_SetApplicationVersion("1.0.0")
}

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	qApp.Exec()
}
