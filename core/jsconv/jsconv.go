// Code generated. DO NOT EDIT.
//
// (C) Copyright Martin Juhlin 2019.
//
// See LICENSE.md for license information

package jsconv

//go:generate go run generator.go

import (
	"runtime"
	"unsafe"

	"github.com/gowebapi/webapi/core/js"
)

// asUint8Array ensures the returned js.Value is a Uint8Array.
// If v is already a Uint8Array, it is returned as is.
// If v is another TypedArray, a Uint8Array view over its underlying buffer is returned.
func asUint8Array(v js.Value) js.Value {
	if v.InstanceOf(js.Global().Get("Uint8Array")) {
		return v
	}
	buf := v.Get("buffer")
	if buf.Truthy() {
		return js.Global().Get("Uint8Array").New(buf, v.Get("byteOffset"), v.Get("byteLength"))
	}
	return v
}

// UInt8ToJs is converting []uint8 to a new javascript Uint8Array.
func UInt8ToJs(src []uint8) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src))
	js.CopyBytesToJS(array, src)
	return array
}

// JsToUInt8 convert javascript Uint8Array to Go uint8.
func JsToUInt8(array js.Value) []uint8 {
	size := array.Get("length").Int()
	mem := make([]byte, size)
	js.CopyBytesToGo(mem, array)
	return mem
}

// CopyToJSUInt8 copies elements from src into existing JavaScript Uint8Array dst.
// Returns the number of elements copied.
func CopyToJSUInt8(dst js.Value, src []uint8) int {
	return js.CopyBytesToJS(dst, src)
}

// CopyFromJSUInt8 copies elements from JavaScript Uint8Array src into Go slice dst.
// Returns the number of elements copied.
func CopyFromJSUInt8(dst []uint8, src js.Value) int {
	return js.CopyBytesToGo(dst, src)
}

// Int8ToJs is converting []int8 to a new javascript Int8Array.
func Int8ToJs(src []int8) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Int8Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 1)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*1)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Int8Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToInt8 convert javascript Int8Array to Go int8.
func JsToInt8(array js.Value) []int8 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*1)

	dst := make([]int8, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*1)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSInt8 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Int8Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSInt8(dst js.Value, src []int8) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*1)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 1
}

// CopyFromJSInt8 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Int8Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSInt8(dst []int8, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*1)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 1
}

// UInt16ToJs is converting []uint16 to a new javascript Uint16Array.
func UInt16ToJs(src []uint16) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Uint16Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 2)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*2)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Uint16Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToUInt16 convert javascript Uint16Array to Go uint16.
func JsToUInt16(array js.Value) []uint16 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*2)

	dst := make([]uint16, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*2)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSUInt16 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Uint16Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSUInt16(dst js.Value, src []uint16) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*2)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 2
}

// CopyFromJSUInt16 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Uint16Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSUInt16(dst []uint16, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*2)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 2
}

// Int16ToJs is converting []int16 to a new javascript Int16Array.
func Int16ToJs(src []int16) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Int16Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 2)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*2)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Int16Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToInt16 convert javascript Int16Array to Go int16.
func JsToInt16(array js.Value) []int16 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*2)

	dst := make([]int16, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*2)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSInt16 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Int16Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSInt16(dst js.Value, src []int16) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*2)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 2
}

// CopyFromJSInt16 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Int16Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSInt16(dst []int16, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*2)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 2
}

// UInt32ToJs is converting []uint32 to a new javascript Uint32Array.
func UInt32ToJs(src []uint32) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Uint32Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 4)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Uint32Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToUInt32 convert javascript Uint32Array to Go uint32.
func JsToUInt32(array js.Value) []uint32 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*4)

	dst := make([]uint32, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*4)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSUInt32 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Uint32Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSUInt32(dst js.Value, src []uint32) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 4
}

// CopyFromJSUInt32 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Uint32Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSUInt32(dst []uint32, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*4)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 4
}

// Int32ToJs is converting []int32 to a new javascript Int32Array.
func Int32ToJs(src []int32) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Int32Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 4)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Int32Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToInt32 convert javascript Int32Array to Go int32.
func JsToInt32(array js.Value) []int32 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*4)

	dst := make([]int32, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*4)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSInt32 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Int32Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSInt32(dst js.Value, src []int32) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 4
}

// CopyFromJSInt32 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Int32Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSInt32(dst []int32, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*4)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 4
}

// Float32ToJs is converting []float32 to a new javascript Float32Array.
func Float32ToJs(src []float32) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Float32Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 4)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Float32Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToFloat32 convert javascript Float32Array to Go float32.
func JsToFloat32(array js.Value) []float32 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*4)

	dst := make([]float32, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*4)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSFloat32 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Float32Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSFloat32(dst js.Value, src []float32) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 4
}

// CopyFromJSFloat32 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Float32Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSFloat32(dst []float32, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*4)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 4
}

// Float64ToJs is converting []float64 to a new javascript Float64Array.
func Float64ToJs(src []float64) js.Value {
	if len(src) == 0 {
		return js.Global().Get("Float64Array").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * 8)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*8)
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("Float64Array").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsToFloat64 convert javascript Float64Array to Go float64.
func JsToFloat64(array js.Value) []float64 {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*8)

	dst := make([]float64, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*8)
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJSFloat64 copies elements from src into existing JavaScript TypedArray dst.
// dst can be a Float64Array or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJSFloat64(dst js.Value, src []float64) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*8)
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / 8
}

// CopyFromJSFloat64 copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a Float64Array or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJSFloat64(dst []float64, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*8)
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / 8
}
