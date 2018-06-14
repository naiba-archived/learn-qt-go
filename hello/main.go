package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)
	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))
	gui.QGuiApplication_Exec()
}
