package webapi_test

import (
	"fmt"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html"
)

func Example_buttonTest() {
	element := webapi.GetWindow().Document().GetElementById("foo")
	button := html.HTMLButtonElementFromJS(element)
	button.SetInnerText("Press me!")

	count := 1
	callback := domcore.EventHandlerToJS(func(event *domcore.Event) interface{} {
		button.SetInnerText(fmt.Sprint("Count: ", count))
		count++
		return nil
	})
	button.SetOnclick(callback)

	// prevent to program to terminate
	// c := make(chan struct{}, 0)
	// <-c
}
