package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//MyItem my item
type MyItem struct {
	_ func() `constructor:"init"`
	widgets.QGraphicsItem
}

func (m *MyItem) init() {
	m.ConnectBoundingRect(m.boundingRect)
	m.ConnectPaint(m.paint)
}

func (m *MyItem) boundingRect() *core.QRectF {
	penWidth := 1.00
	return core.NewQRectF4(0-penWidth/2, 0-penWidth/2,
		20+penWidth, 20+penWidth)
}

func (m *MyItem) paint(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget) {
	painter.DrawRect2(0, 0, 20, 20)
}
