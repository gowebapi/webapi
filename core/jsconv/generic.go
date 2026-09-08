package jsconv

import (
	"github.com/gowebapi/webapi/core/js"
)

// Supported represents the scalar numeric types supported for TypedArray conversions.
type Supported interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~float32 | ~float64
}

// ToJS converts a Go numeric slice to a newly allocated JavaScript TypedArray.
func ToJS[E Supported](src []E) js.Value {
	switch s := any(src).(type) {
	case []uint8:
		return UInt8ToJs(s)
	case []int8:
		return Int8ToJs(s)
	case []uint16:
		return UInt16ToJs(s)
	case []int16:
		return Int16ToJs(s)
	case []uint32:
		return UInt32ToJs(s)
	case []int32:
		return Int32ToJs(s)
	case []float32:
		return Float32ToJs(s)
	case []float64:
		return Float64ToJs(s)
	default:
		panic("unsupported type")
	}
}

// FromJS converts a JavaScript TypedArray to a newly allocated Go slice.
func FromJS[E Supported](jsArray js.Value) []E {
	var zero E
	switch any(zero).(type) {
	case uint8:
		return any(JsToUInt8(jsArray)).([]E)
	case int8:
		return any(JsToInt8(jsArray)).([]E)
	case uint16:
		return any(JsToUInt16(jsArray)).([]E)
	case int16:
		return any(JsToInt16(jsArray)).([]E)
	case uint32:
		return any(JsToUInt32(jsArray)).([]E)
	case int32:
		return any(JsToInt32(jsArray)).([]E)
	case float32:
		return any(JsToFloat32(jsArray)).([]E)
	case float64:
		return any(JsToFloat64(jsArray)).([]E)
	default:
		panic("unsupported type")
	}
}

// CopyToJS copies elements from Go slice src into existing JavaScript TypedArray dst.
// dst can be a matching TypedArray or a Uint8Array.
// Returns the number of elements (not bytes) copied.
func CopyToJS[E Supported](dst js.Value, src []E) int {
	switch s := any(src).(type) {
	case []uint8:
		return CopyToJSUInt8(dst, s)
	case []int8:
		return CopyToJSInt8(dst, s)
	case []uint16:
		return CopyToJSUInt16(dst, s)
	case []int16:
		return CopyToJSInt16(dst, s)
	case []uint32:
		return CopyToJSUInt32(dst, s)
	case []int32:
		return CopyToJSInt32(dst, s)
	case []float32:
		return CopyToJSFloat32(dst, s)
	case []float64:
		return CopyToJSFloat64(dst, s)
	default:
		panic("unsupported type")
	}
}

// CopyFromJS copies elements from JavaScript TypedArray src into existing Go slice dst.
// src can be a matching TypedArray or a Uint8Array.
// Returns the number of elements (not bytes) copied.
func CopyFromJS[E Supported](dst []E, src js.Value) int {
	switch d := any(dst).(type) {
	case []uint8:
		return CopyFromJSUInt8(d, src)
	case []int8:
		return CopyFromJSInt8(d, src)
	case []uint16:
		return CopyFromJSUInt16(d, src)
	case []int16:
		return CopyFromJSInt16(d, src)
	case []uint32:
		return CopyFromJSUInt32(d, src)
	case []int32:
		return CopyFromJSInt32(d, src)
	case []float32:
		return CopyFromJSFloat32(d, src)
	case []float64:
		return CopyFromJSFloat64(d, src)
	default:
		panic("unsupported type")
	}
}
