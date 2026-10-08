//go:build js && wasm

package webgl_test

import (
	"testing"
	"unsafe"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/core/js"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/graphics/webgl"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/support/humanize"
)

var webglSizes = []struct {
	floatCount int
}{
	{1024},            // 4 KiB
	{16 * 1024},       // 64 KiB
	{256 * 1024},      // 1 MiB
	{2 * 1024 * 1024}, // 8 MiB
}

func initGL(t testing.TB) (*webgl.RenderingContext, *webgl.Buffer) {
	doc := webapi.GetWindow().Document()
	if doc == nil || doc.JSValue().IsNull() || doc.JSValue().IsUndefined() {
		t.Skip("window.document not available")
	}

	canvasElem := doc.CreateElement("canvas", nil)
	if canvasElem == nil {
		t.Skip("failed to create canvas element")
	}
	canvasHTML := canvas.HTMLCanvasElementFromWrapper(canvasElem)
	canvasHTML.SetWidth(800)
	canvasHTML.SetHeight(600)

	opts := js.Global().Get("Object").New()
	opts.Set("failIfMajorPerformanceCaveat", false)

	ctxObj := canvasHTML.GetContext("webgl", opts)
	if ctxObj == nil || ctxObj.JSValue().IsNull() || ctxObj.JSValue().IsUndefined() {
		// Try experimental-webgl
		ctxObj = canvasHTML.GetContext("experimental-webgl", opts)
		if ctxObj == nil || ctxObj.JSValue().IsNull() || ctxObj.JSValue().IsUndefined() {
			t.Skip("WebGL not supported in this browser environment (try running with WASM_HEADLESS=off)")
		}
	}

	gl := webgl.RenderingContextFromWrapper(ctxObj)
	if gl == nil || gl.JSValue().IsNull() {
		t.Skip("failed to obtain WebGL rendering context")
	}

	buf := gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, buf)
	return gl, buf
}

// BenchmarkWebGL_AllocCopy_BufferData measures allocating a new Float32Array on each call and uploading with bufferData.
func BenchmarkWebGL_AllocCopy_BufferData(b *testing.B) {
	gl, _ := initGL(b)

	for _, tc := range webglSizes {
		byteSize := tc.floatCount * 4
		b.Run(humanize.Bytes(byteSize), func(b *testing.B) {
			data := make([]float32, tc.floatCount)
			for i := range data {
				data[i] = float32(i) * 0.1
			}
			b.SetBytes(int64(byteSize))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				f32JS := jsconv.Float32ToJs(data)
				gl.BufferData2(webgl.ARRAY_BUFFER, webgl.UnionFromJS(f32JS), webgl.DYNAMIC_DRAW)
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkWebGL_ReusedCopy_BufferSubData measures copying into a preallocated Float32Array and uploading via bufferSubData.
func BenchmarkWebGL_ReusedCopy_BufferSubData(b *testing.B) {
	gl, _ := initGL(b)

	for _, tc := range webglSizes {
		byteSize := tc.floatCount * 4
		b.Run(humanize.Bytes(byteSize), func(b *testing.B) {
			data := make([]float32, tc.floatCount)
			for i := range data {
				data[i] = float32(i) * 0.1
			}
			// Preallocate GL buffer storage
			gl.BufferData(webgl.ARRAY_BUFFER, byteSize, webgl.DYNAMIC_DRAW)

			// Preallocate JS Float32Array
			dstJS := js.Global().Get("Float32Array").New(tc.floatCount)
			dstUnion := webgl.UnionFromJS(dstJS)

			b.SetBytes(int64(byteSize))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				jsconv.CopyToJSFloat32(dstJS, data)
				gl.BufferSubData(webgl.ARRAY_BUFFER, 0, dstUnion)
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkWebGL_ReusedCopy_BufferData measures copying into a preallocated Float32Array and reallocating GPU buffer via bufferData.
func BenchmarkWebGL_ReusedCopy_BufferData(b *testing.B) {
	gl, _ := initGL(b)

	for _, tc := range webglSizes {
		byteSize := tc.floatCount * 4
		b.Run(humanize.Bytes(byteSize), func(b *testing.B) {
			data := make([]float32, tc.floatCount)
			for i := range data {
				data[i] = float32(i) * 0.1
			}
			dstJS := js.Global().Get("Float32Array").New(tc.floatCount)
			dstUnion := webgl.UnionFromJS(dstJS)

			b.SetBytes(int64(byteSize))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				jsconv.CopyToJSFloat32(dstJS, data)
				gl.BufferData2(webgl.ARRAY_BUFFER, dstUnion, webgl.DYNAMIC_DRAW)
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkWebGL_DirectArrayBufferView_BufferSubData measures creating a TypedArray view directly over an existing ArrayBuffer (simulating zero-copy memory view) and uploading via bufferSubData.
func BenchmarkWebGL_DirectArrayBufferView_BufferSubData(b *testing.B) {
	gl, _ := initGL(b)

	for _, tc := range webglSizes {
		byteSize := tc.floatCount * 4
		b.Run(humanize.Bytes(byteSize), func(b *testing.B) {
			data := make([]float32, tc.floatCount)
			// Preallocate GL buffer
			gl.BufferData(webgl.ARRAY_BUFFER, byteSize, webgl.DYNAMIC_DRAW)

			// Allocate a standalone ArrayBuffer representing external/linear memory
			arrayBuf := js.Global().Get("ArrayBuffer").New(byteSize)
			f32Ctor := js.Global().Get("Float32Array")

			b.SetBytes(int64(byteSize))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// Zero-copy: create Float32Array view at offset 0 of arrayBuf without copying bytes
				ptr := uintptr(unsafe.Pointer(&data[0]))
				_ = ptr
				view := f32Ctor.New(arrayBuf, 0, tc.floatCount)
				gl.BufferSubData(webgl.ARRAY_BUFFER, 0, webgl.UnionFromJS(view))
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}
