package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

type viewer struct {
	core.QObject

	_ string `property:"imageURL"`
	_ bool `property:"hasError"`
}

var v *viewer

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	initViewer()

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("v", v)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}

func initViewer() {
	v = NewViewer(nil)
	v.SetImageURL("test")
}
