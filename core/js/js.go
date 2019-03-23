// Part of this is taken from Go language source code:
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

// Package js is equal api to syscall/js but is panic on all method invocation
// when compiled under none wasm taget.
//
// The usage is to get a tab complession to work inside an IDE
// (e.g. Visual Studio Code) without changing to GOOS to js
package js

const message = "this library should be executed under WASM"

// Error wraps a JavaScript error.
type Error struct {
	// Value is the underlying JavaScript error value.
	Value
}

// Error implements the error interface.
func (e Error) Error() string {
	panic(message)
}

// Func is a wrapped Go function to be called by JavaScript.
type Func struct {
	Value
}

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
	panic(message)
}

// Release frees up resources allocated for the function.
// The function must not be invoked after calling Release.
func (c Func) Release() {
	panic(message)
}

// Type represents the JavaScript type of a Value.
type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string {
	panic(message)
}

// TypedArray represents a JavaScript typed array.
type TypedArray struct {
	Value
}

// TypedArrayOf returns a JavaScript typed array backed by the slice's underlying array.
//
// The supported types are []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
// Passing an unsupported value causes a panic.
//
// TypedArray.Release must be called to free up resources when the typed array will not be used any more.
func TypedArrayOf(slice interface{}) TypedArray {
	panic(message)
}

// Release frees up resources allocated for the typed array.
// The typed array and its buffer must not be accessed after calling Release.
func (a TypedArray) Release() {
	panic(message)
}

// Value represents a JavaScript value. The zero value is the JavaScript value "undefined".
type Value struct {
}

// Global returns the JavaScript global object, usually "window" or "global".
func Global() Value {
	panic(message)
}

// Null returns the JavaScript value "null".
func Null() Value {
	panic(message)
}

// Undefined returns the JavaScript value "undefined".
func Undefined() Value {
	panic(message)
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
	panic(message)
}

// Bool returns the value v as a bool. It panics if v is not a JavaScript boolean.
func (v Value) Bool() bool {
	panic(message)
}

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Call(m string, args ...interface{}) Value {
	panic(message)
}

// Float returns the value v as a float64. It panics if v is not a JavaScript number.
func (v Value) Float() float64 {
	panic(message)
}

// Get returns the JavaScript property p of value v.
func (v Value) Get(p string) Value {
	panic(message)
}

// Index returns JavaScript index i of value v.
func (v Value) Index(i int) Value {
	panic(message)
}

// InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.
func (v Value) InstanceOf(t Value) bool {
	panic(message)
}

// Int returns the value v truncated to an int. It panics if v is not a JavaScript number.
func (v Value) Int() int {
	panic(message)
}

// Invoke does a JavaScript call of the value v with the given arguments.
// It panics if v is not a function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Invoke(args ...interface{}) Value {
	panic(message)
}

// JSValue implements Wrapper interface.
func (v Value) JSValue() Value {
	panic(message)
}

// Length returns the JavaScript property "length" of v.
func (v Value) Length() int {
	panic(message)
}

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) New(args ...interface{}) Value {
	panic(message)
}

// Set sets the JavaScript property p of value v to ValueOf(x).
func (v Value) Set(p string, x interface{}) {
	panic(message)
}

// SetIndex sets the JavaScript index i of value v to ValueOf(x).
func (v Value) SetIndex(i int, x interface{}) {
	panic(message)
}

// String returns the value v converted to string according to JavaScript type conversions.
func (v Value) String() string {
	panic(message)
}

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (v Value) Truthy() bool {
	panic(message)
}

// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
// except that it returns TypeNull instead of TypeObject for null.
func (v Value) Type() Type {
	panic(message)
}

// A ValueError occurs when a Value method is invoked on
// a Value that does not support it. Such cases are documented
// in the description of each method.
type ValueError struct {
	Method string
	Type   Type
}

func (e *ValueError) Error() string {
	panic(message)
}

// Wrapper is implemented by types that are backed by a JavaScript value.
type Wrapper interface {
	// JSValue returns a JavaScript value associated with an object.
	JSValue() Value
}
