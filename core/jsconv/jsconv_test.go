//go:build js && wasm

package jsconv_test

import (
	"math"
	"testing"

	"github.com/gowebapi/webapi/core/js"
	"github.com/gowebapi/webapi/core/jsconv"
)

func TestGenericsAndCopies(t *testing.T) {
	// Test Float32 ToJS / FromJS
	f32 := []float32{1.5, 2.25, -3.125, 4.0}
	jsF32 := jsconv.ToJS(f32)
	if jsF32.Type() != js.TypeObject {
		t.Fatalf("expected js object, got %v", jsF32.Type())
	}
	f32RoundTrip := jsconv.FromJS[float32](jsF32)
	if len(f32RoundTrip) != len(f32) {
		t.Fatalf("expected length %d, got %d", len(f32), len(f32RoundTrip))
	}
	for i := range f32 {
		if f32[i] != f32RoundTrip[i] {
			t.Errorf("f32[%d] = %f, expected %f", i, f32RoundTrip[i], f32[i])
		}
	}

	// Test CopyToJS and CopyFromJS with Float32
	f32Dst := js.Global().Get("Float32Array").New(len(f32))
	n := jsconv.CopyToJS(f32Dst, f32)
	if n != len(f32) {
		t.Fatalf("expected %d elements copied, got %d", len(f32), n)
	}

	f32Back := make([]float32, len(f32))
	nBack := jsconv.CopyFromJS(f32Back, f32Dst)
	if nBack != len(f32) {
		t.Fatalf("expected %d elements copied back, got %d", len(f32), nBack)
	}
	for i := range f32 {
		if f32[i] != f32Back[i] {
			t.Errorf("f32Back[%d] = %f, expected %f", i, f32Back[i], f32[i])
		}
	}

	// Test directly typed CopyToJSFloat32 / CopyFromJSFloat32
	f32New := []float32{10.0, 20.0, 30.0, 40.0}
	nTyped := jsconv.CopyToJSFloat32(f32Dst, f32New)
	if nTyped != len(f32New) {
		t.Fatalf("expected %d elements copied by typed copy, got %d", len(f32New), nTyped)
	}
	f32TypedBack := make([]float32, len(f32New))
	nTypedBack := jsconv.CopyFromJSFloat32(f32TypedBack, f32Dst)
	if nTypedBack != len(f32New) {
		t.Fatalf("expected %d elements read back by typed copy, got %d", len(f32New), nTypedBack)
	}
	for i := range f32New {
		if f32New[i] != f32TypedBack[i] {
			t.Errorf("f32TypedBack[%d] = %f, expected %f", i, f32TypedBack[i], f32New[i])
		}
	}

	// Test UInt8
	u8 := []uint8{0, 42, 128, 255}
	jsU8 := jsconv.ToJS(u8)
	u8Back := jsconv.FromJS[uint8](jsU8)
	for i := range u8 {
		if u8[i] != u8Back[i] {
			t.Errorf("u8[%d] = %d, expected %d", i, u8Back[i], u8[i])
		}
	}

	// Test Float64
	f64 := []float64{math.Pi, math.E, -1.0e10}
	jsF64 := jsconv.ToJS(f64)
	f64Back := jsconv.FromJS[float64](jsF64)
	for i := range f64 {
		if f64[i] != f64Back[i] {
			t.Errorf("f64[%d] = %f, expected %f", i, f64Back[i], f64[i])
		}
	}
}
