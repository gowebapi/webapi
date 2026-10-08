# WebAPI

Comprehensive Go WebAssembly (WASM) bindings for standard browser Web APIs (DOM, HTML, CSS, Fetch, WebGL, WebGPU, WebNN, WebCodecs, etc.).

Rather than handwriting thousands of bindings, `webapi` automatically generates idiomatic, type-safe Go bindings from official **WebIDL** specifications using [`webidl-bind`](file:///home/janpf/Projects/gowebapi/webidl-bind/).

[![PkgGoDev](https://pkg.go.dev/badge/github.com/gowebapi/webapi.svg)](https://pkg.go.dev/github.com/gowebapi/webapi)

---

## Features

- **Standard Compliant**: Generated directly from official W3C and WHATWG WebIDL specifications.
- **Dual-File Compilation**:
  - `*_js.go`: Native browser execution under `GOOS=js GOARCH=wasm` using `syscall/js`.
  - `*.go`: Host stub implementations under `//go:build !js` using [`core/js`](file:///home/janpf/Projects/gowebapi/webapi/core/js), enabling standard host toolchains, IDEs (`gopls`), linters, and `go test` to navigate and typecheck without switching environments.
- **High-Throughput Buffer Conversions**: Optimized array buffer copying between Go and JavaScript in [`core/jsconv`](file:///home/janpf/Projects/gowebapi/webapi/core/jsconv/) (`ToJS`, `FromJS`, `CopyToJS`, `CopyFromJS`).
- **Comprehensive API Coverage**: Includes modern browser standards such as WebGPU, WebNN, WebCodecs, Cookie Store, File System Access, Storage Buckets, View Transitions, and Anchor Positioning.

---

## Installation

```bash
go get github.com/gowebapi/webapi
```

Requires **Go 1.22+**.

---

## Example: Button Counter

The following example binds to a DOM button element, sets its text, and increments a counter each time the button is clicked:

```go
package main

import (
	"fmt"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/html/htmlevent"
)

func main() {
	// Get the element from the DOM
	element := webapi.GetWindow().Document().GetElementById("myButton")

	// Cast it into the correct HTML element class
	button := html.HTMLButtonElementFromJS(element.JSValue())

	// Change text
	button.SetInnerText("Press me!")

	// Register an event callback that displays a click counter
	count := 1
	button.SetOnClick(func(event *htmlevent.MouseEvent, currentTarget *html.HTMLElement) {
		button.SetInnerText(fmt.Sprint("Count: ", count))
		count++
	})

	// Prevent program from terminating
	c := make(chan struct{}, 0)
	<-c
}
```

---

## Building and Running

### 1. Compile to WebAssembly

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

### 2. Run in the Browser

Serve `main.wasm` alongside Go's WebAssembly support file (`wasm_exec.js`) and an HTML loader:

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8"/>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</head>
<body>
    <button id="myButton">Loading...</button>
</body>
</html>
```

*(You can copy `wasm_exec.js` from your Go installation via `cp $(go env GOROOT)/misc/wasm/wasm_exec.js .` or `$(go env GOROOT)/lib/wasm/wasm_exec.js` depending on your Go version).*

---

## Testing

- **Host Stub Verification**:
  ```bash
  go test ./...
  ```
- **In-Browser WebAssembly Tests** (using [`wasmbrowsertest`](https://github.com/agnivade/wasmbrowsertest)):
  ```bash
  env -i PATH="$PATH" HOME="$HOME" DISPLAY="${DISPLAY:-:0}" \
    GOOS=js GOARCH=wasm go test -v \
    -exec="wasmbrowsertest" github.com/gowebapi/webapi/dom
  ```

---

## Package Overview

| Domain | Key Packages | Description |
| :--- | :--- | :--- |
| **Root & Core** | [`webapi`](file:///home/janpf/Projects/gowebapi/webapi/webapi.go), [`core`](file:///home/janpf/Projects/gowebapi/webapi/core/) | Global scope accessors (`GetWindow()`, `GetDocument()`), JS interop wrappers |
| **DOM** | `dom`, `dom/domcore`, `dom/geometry`, `dom/eyedropper` | Nodes, Elements, Events, Mutation Observers, Geometry |
| **HTML** | `html`, `html/canvas`, `html/media`, `html/worker` | HTML elements, Canvas 2D, Audio/Video, Web Workers |
| **CSS** | `css`, `css/cssom`, `css/viewtransitions`, `css/anchorposition` | CSSOM, Animations, View Transitions, Anchor Positioning |
| **Graphics & ML** | `graphics/webgl`, `graphics/webgpu`, `ml/webnn` | WebGL 1 & 2, WebGPU, Web Neural Network API |
| **Media** | `media/audio`, `media/webrtc`, `media/webcodecs` | Web Audio, WebRTC, WebCodecs audio/video encoding & decoding |
| **Storage & File** | `storage`, `cookie`, `file/fs`, `storage/buckets`, `storage/weblocks` | Cookie Store, File System Access, Storage Buckets, Web Locks |
| **Device & Hardware** | `device/serial`, `device/battery`, `device/computepressure`, `device/gamepad` | Web Serial, Compute Pressure, Gamepad, Sensor APIs |

---

## Changelog

- 2026/10/08: Modernized Go; Updated with recent APIs (Cookie Store, WebNN, etc.).
