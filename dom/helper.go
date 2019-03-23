package dom

import js "github.com/gowebapi/webapi/core/js"

// RequestFullscreenByBrowser is an alternative way to
// request fullscreen that is taking different browser into
// accound
func (t *Element) RequestFullscreenByBrowser() {
	elem := t.Value_JS
	if elem.Get("requestFullscreen") != js.Undefined() {
		elem.Call("requestFullscreen", "")
	} else if elem.Get("mozRequestFullScreen") != js.Undefined() { /* Firefox */
		elem.Call("mozRequestFullScreen", "")
	} else if elem.Get("webkitRequestFullscreen") != js.Undefined() { /* Chrome, Safari and Opera */
		elem.Call("webkitRequestFullscreen", "")
	} else if elem.Get("msRequestFullscreen") != js.Undefined() { /* IE/Edge */
		elem.Call("msRequestFullscreen", "")
	}
}
