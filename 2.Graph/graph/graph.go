package main

import (
	"github.com/therecipe/qt/widgets"
)

//GraphWindow graph window
type GraphWindow struct {
	widgets.QMainWindow

	view  *widgets.QGraphicsView
	scene *widgets.QGraphicsScene
}

func initGraph() *GraphWindow {
	var this = NewGraphWindow(nil, 0)
	this.view = widgets.NewQGraphicsView(this)
	this.scene = widgets.NewQGraphicsScene(this)

	this.SetWindowTitle("图形视图框架")
	this.SetFixedSize(this.Size())
	this.view.SetScene(this.scene)
	this.view.SetFixedSize(this.Size())

	item := widgets.NewQGraphicsRectItem3(100, 100, 50, 50, nil)
	this.scene.AddItem(item)

	return this
}
