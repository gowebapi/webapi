package view_test

import (
	"github.com/gowebapi/webapi/css/cssom/view"
	"github.com/gowebapi/webapi/dom"
)

func ExampleCaretPosition_OffsetNode() {
	var caret *view.CaretPosition

	// cast to correct type
	value := dom.NodeFromJS(caret.OffsetNode())

	// do something with value
	_ = value
}
