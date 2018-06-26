package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//NewMyItem my item
func NewMyItem() *widgets.QGraphicsItem {
	this := widgets.NewQGraphicsItem(nil)
	this.ConnectBoundingRect(func() *core.QRectF {
		penWidth := 1.00
		return core.NewQRectF4(0-penWidth/2, 0-penWidth/2,
			20+penWidth, 20+penWidth)
	})
	this.ConnectPaint(func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget) {
		brush := gui.NewQBrush()
		brush.SetColor2(core.Qt__red)
		painter.SetBrush(brush)
		painter.DrawRect2(0, 0, 100, 100)
	})

	this.SetToolTip("test tooltip")

	return this
}
