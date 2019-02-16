package webapi_test

import (
	"fmt"

	"github.com/gowebapi/webapi"
	js "github.com/gowebapi/webapi/core/failjs"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html"
)

func Example_buttonTest() {
	element := webapi.GetWindow().Document().GetElementById("foo")
	button := html.HTMLButtonElementFromJS(element.JSValue())
	button.SetInnerText("Press me!")

	count := 1
	callback := domcore.EventHandlerToJS(func(event *domcore.Event) js.Value {
		button.SetInnerText(fmt.Sprint("Count: ", count))
		count++
		return js.Null()
	})
	button.SetOnclick(callback)

	// prevent to program to terminate
	// c := make(chan struct{}, 0)
	// <-c
}
