package dom

import js "github.com/gowebapi/webapi/core/js"

// RequestFullscreenByBrowser is an alternative way to
// request fullscreen that is taking different browser into
// accound
func (t *Element) RequestFullscreenByBrowser() {
	elem := t.Value_JS
	if !elem.Get("requestFullscreen").IsUndefined() {
		elem.Call("requestFullscreen", "")
	} else if !elem.Get("mozRequestFullScreen").IsUndefined() { /* Firefox */
		elem.Call("mozRequestFullScreen", "")
	} else if !elem.Get("webkitRequestFullscreen").IsUndefined() { /* Chrome, Safari and Opera */
		elem.Call("webkitRequestFullscreen", "")
	} else if !elem.Get("msRequestFullscreen").IsUndefined() { /* IE/Edge */
		elem.Call("msRequestFullscreen", "")
	}
}
