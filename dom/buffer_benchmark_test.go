//go:build js && wasm

package dom_test

import (
	"fmt"
	"testing"

	"github.com/gowebapi/webapi/core/js"
	"github.com/gowebapi/webapi/core/jsconv"
	"github.com/gowebapi/webapi/support/humanize"
)

var benchmarkSizes = []int{
	1024,
	64 * 1024,
	1024 * 1024,
	8 * 1024 * 1024,
}

// BenchmarkDOMBuffer_AllocCopy measures allocating a new Uint8Array in JS and copying bytes into it.
func BenchmarkDOMBuffer_AllocCopy(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(humanize.Bytes(size), func(b *testing.B) {
			data := make([]byte, size)
			for i := range data {
				data[i] = byte(i)
			}
			b.SetBytes(int64(size))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				u8 := jsconv.UInt8ToJs(data)
				_ = u8
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkDOMBuffer_ReusedCopy measures copying bytes into an existing pre-allocated Uint8Array.
func BenchmarkDOMBuffer_ReusedCopy(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(humanize.Bytes(size), func(b *testing.B) {
			data := make([]byte, size)
			for i := range data {
				data[i] = byte(i)
			}
			dst := js.Global().Get("Uint8Array").New(size)
			b.SetBytes(int64(size))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				n := jsconv.CopyToJSUInt8(dst, data)
				if n != size {
					b.Fatalf("expected %d copied, got %d", size, n)
				}
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkDOMBuffer_RoundTrip measures Go -> JS -> Go buffer copy.
func BenchmarkDOMBuffer_RoundTrip(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(humanize.Bytes(size), func(b *testing.B) {
			src := make([]byte, size)
			dst := make([]byte, size)
			for i := range src {
				src[i] = byte(i)
			}
			jsBuf := js.Global().Get("Uint8Array").New(size)
			b.SetBytes(int64(size * 2)) // Go->JS + JS->Go
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				jsconv.CopyToJSUInt8(jsBuf, src)
				jsconv.CopyFromJSUInt8(dst, jsBuf)
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkDOMBuffer_ToBlob measures transferring a buffer from Go to JS and wrapping it in a DOM Blob.
func BenchmarkDOMBuffer_ToBlob(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(humanize.Bytes(size), func(b *testing.B) {
			data := make([]byte, size)
			b.SetBytes(int64(size))
			blobCtor := js.Global().Get("Blob")
			arrayCtor := js.Global().Get("Array")
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				u8 := jsconv.UInt8ToJs(data)
				parts := arrayCtor.New(1)
				parts.SetIndex(0, u8)
				blob := blobCtor.New(parts)
				_ = blob
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}

// BenchmarkDOMBuffer_ElementDataset measures transferring a buffer to JS and storing it on a DOM Element.
func BenchmarkDOMBuffer_ElementDataset(b *testing.B) {
	doc := js.Global().Get("document")
	if doc.IsUndefined() || doc.IsNull() {
		b.Skip("document not available")
	}
	elem := doc.Call("createElement", "div")
	doc.Get("body").Call("appendChild", elem)
	defer doc.Get("body").Call("removeChild", elem)

	for _, size := range benchmarkSizes {
		b.Run(humanize.Bytes(size), func(b *testing.B) {
			data := make([]byte, size)
			dst := js.Global().Get("Uint8Array").New(size)
			elem.Set(fmt.Sprintf("_buf_%s", humanize.Bytes(size)), dst)
			b.SetBytes(int64(size))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				jsconv.CopyToJSUInt8(dst, data)
			}

			elapsed := b.Elapsed()
			if elapsed > 0 && b.N > 0 {
				b.ReportMetric(humanize.DurationPerOp(elapsed, b.N))
			}
		})
	}
}
