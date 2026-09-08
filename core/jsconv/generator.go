//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"
)

const templateText = `
{{define "file-header"}}
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

{{end}}

{{define "type-uint8"}}

// {{.Name}}ToJs is converting []{{.GoType}} to a new javascript {{.Js}}.
func {{.Name}}ToJs(src []{{.GoType}}) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src))
	js.CopyBytesToJS(array, src)
	return array
}

// JsTo{{.Name}} convert javascript {{.Js}} to Go {{.GoType}}.
func JsTo{{.Name}}(array js.Value) []{{.GoType}} {
	size := array.Get("length").Int()
	mem := make([]byte, size)
	js.CopyBytesToGo(mem, array)
	return mem
}

// CopyToJS{{.Name}} copies elements from src into existing JavaScript {{.Js}} dst.
// Returns the number of elements copied.
func CopyToJS{{.Name}}(dst js.Value, src []{{.GoType}}) int {
	return js.CopyBytesToJS(dst, src)
}

// CopyFromJS{{.Name}} copies elements from JavaScript {{.Js}} src into Go slice dst.
// Returns the number of elements copied.
func CopyFromJS{{.Name}}(dst []{{.GoType}}, src js.Value) int {
	return js.CopyBytesToGo(dst, src)
}

{{end}}

{{define "type"}}

// {{.Name}}ToJs is converting []{{.GoType}} to a new javascript {{.Js}}.
func {{.Name}}ToJs(src []{{.GoType}}) js.Value {
	if len(src) == 0 {
		return js.Global().Get("{{.Js}}").New(0)
	}
	array := js.Global().Get("Uint8Array").New(len(src) * {{.Size}})
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*{{.Size}})
	js.CopyBytesToJS(array, raw)
	runtime.KeepAlive(src)

	buf := array.Get("buffer")
	return js.Global().Get("{{.Js}}").New(
		buf,
		array.Get("byteOffset"),
		len(src),
	)
}

// JsTo{{.Name}} convert javascript {{.Js}} to Go {{.GoType}}.
func JsTo{{.Name}}(array js.Value) []{{.GoType}} {
	size := array.Get("length").Int()
	if size == 0 {
		return nil
	}
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf, array.Get("byteOffset"), size*{{.Size}})

	dst := make([]{{.GoType}}, size)
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), size*{{.Size}})
	js.CopyBytesToGo(raw, tmp)
	runtime.KeepAlive(dst)
	return dst
}

// CopyToJS{{.Name}} copies elements from src into existing JavaScript TypedArray dst.
// dst can be a {{.Js}} or a Uint8Array of sufficient capacity.
// Returns the number of elements copied.
func CopyToJS{{.Name}}(dst js.Value, src []{{.GoType}}) int {
	if len(src) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*{{.Size}})
	u8 := asUint8Array(dst)
	n := js.CopyBytesToJS(u8, raw)
	runtime.KeepAlive(src)
	return n / {{.Size}}
}

// CopyFromJS{{.Name}} copies elements from JavaScript TypedArray src into Go slice dst.
// src can be a {{.Js}} or a Uint8Array.
// Returns the number of elements copied.
func CopyFromJS{{.Name}}(dst []{{.GoType}}, src js.Value) int {
	if len(dst) == 0 {
		return 0
	}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(&dst[0])), len(dst)*{{.Size}})
	u8 := asUint8Array(src)
	n := js.CopyBytesToGo(raw, u8)
	runtime.KeepAlive(dst)
	return n / {{.Size}}
}

{{end}}
`

var tmpl = template.Must(template.New("main").Parse(templateText))

type TypeInfo struct {
	GoType string
	Name   string
	Js     string
	Size   int
}

var types = []TypeInfo{
	{"uint8", "UInt8", "Uint8Array", 1},
	{"int8", "Int8", "Int8Array", 1},
	{"uint16", "UInt16", "Uint16Array", 2},
	{"int16", "Int16", "Int16Array", 2},
	{"uint32", "UInt32", "Uint32Array", 4},
	{"int32", "Int32", "Int32Array", 4},
	{"float32", "Float32", "Float32Array", 4},
	{"float64", "Float64", "Float64Array", 8},
}

func main() {
	content, err := generateText()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	if source, err := format.Source(content); err == nil {
		content = source
	} else {
		fmt.Fprintf(os.Stderr, "error:unable to format output source code: %s\n", err)
	}
	if err := os.WriteFile("jsconv.go", content, 0664); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func generateText() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "file-header", nil); err != nil {
		return nil, err
	}
	for _, q := range types {
		if t := tmpl.Lookup("type-" + q.GoType); t != nil {
			if err := t.Execute(buf, q); err != nil {
				return nil, err
			}
		} else if err := tmpl.ExecuteTemplate(buf, "type", q); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
