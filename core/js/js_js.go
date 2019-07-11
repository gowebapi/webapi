// Part of this is taken from Go language source code:
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package js is equal api to syscall/js but is panic on all method invocation
// when compiled under none wasm taget.
//
// The usage is to get a tab complession to work inside an IDE
// (e.g. Visual Studio Code) without changing to GOOS to js
package js

import (
	"syscall/js"
)

// Error wraps a JavaScript error.
type Error = js.Error

// Func is a wrapped Go function to be called by JavaScript.
type Func = js.Func

// FuncOf returns a wrapped function.
//
// Invoking the JavaScript function will synchronously call the Go function fn with the value of JavaScript's
// "this" keyword and the arguments of the invocation.
// The return value of the invocation is the result of the Go function mapped back to JavaScript according to ValueOf.
//
// A wrapped function triggered during a call from Go to JavaScript gets executed on the same goroutine.
// A wrapped function triggered by JavaScript's event loop gets executed on an extra goroutine.
// Blocking operations in the wrapped function will block the event loop.
// As a consequence, if one wrapped function blocks, other wrapped funcs will not be processed.
// A blocking function should therefore explicitly start a new goroutine.
//
// Func.Release must be called to free up resources when the function will not be used any more.
func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return js.FuncOf(fn)
}

// Type represents the JavaScript type of a Value.
type Type = js.Type

const (
	TypeUndefined Type = js.TypeUndefined
	TypeNull           = js.TypeNull
	TypeBoolean        = js.TypeBoolean
	TypeNumber         = js.TypeNumber
	TypeString         = js.TypeString
	TypeSymbol         = js.TypeSymbol
	TypeObject         = js.TypeObject
	TypeFunction       = js.TypeFunction
)

// TypedArray represents a JavaScript typed array.
type TypedArray = js.TypedArray

// TypedArrayOf returns a JavaScript typed array backed by the slice's underlying array.
//
// The supported types are []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
// Passing an unsupported value causes a panic.
//
// TypedArray.Release must be called to free up resources when the typed array will not be used any more.
func TypedArrayOf(slice interface{}) TypedArray {
	return TypedArray(js.TypedArrayOf(slice))
}

// Value represents a JavaScript value. The zero value is the JavaScript value "undefined".
type Value = js.Value

// Global returns the JavaScript global object, usually "window" or "global".
func Global() Value {
	return js.Global()
}

// Null returns the JavaScript value "null".
func Null() Value {
	return js.Null()
}

// Undefined returns the JavaScript value "undefined".
func Undefined() Value {
	return js.Undefined()
}

// ValueOf returns x as a JavaScript value:
//
//  | Go                     | JavaScript             |
//  | ---------------------- | ---------------------- |
//  | js.Value               | [its value]            |
//  | js.TypedArray          | typed array            |
//  | js.Func                | function               |
//  | nil                    | null                   |
//  | bool                   | boolean                |
//  | integers and floats    | number                 |
//  | string                 | string                 |
//  | []interface{}          | new array              |
//  | map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.
func ValueOf(x interface{}) Value {
	return js.ValueOf(x)
}

// A ValueError occurs when a Value method is invoked on
// a Value that does not support it. Such cases are documented
// in the description of each method.
type ValueError = js.ValueError

// Wrapper is implemented by types that are backed by a JavaScript value.
type Wrapper = js.Wrapper
