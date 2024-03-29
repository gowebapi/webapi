// Part of this is taken from Go language source code:
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package js is the same api as syscall/js, but when not
// compiling under WASM target it panics.
//
// The objective is to get tab completion to work inside an IDE
// (e.g. Visual Studio Code, GoLand) without having to change GOOS to js.
//
// Original documentation:
//
// Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
// Its API is based on JavaScript semantics.
//
// This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a
// comprehensive API for users. It is exempt from the Go compatibility promise.
package js

import (
	"syscall/js"
)

// CopyBytesToGo copies bytes from the Uint8Array src to dst.
// It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
// CopyBytesToGo panics if src is not an Uint8Array.
func CopyBytesToGo(dst []byte, src Value) int {
	return js.CopyBytesToGo(dst, src)
}

// CopyBytesToJS copies bytes from src to the Uint8Array dst.
// It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
// CopyBytesToJS panics if dst is not an Uint8Array.
func CopyBytesToJS(dst Value, src []byte) int {
	return js.CopyBytesToJS(dst, src)
}

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
