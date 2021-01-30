// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
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
	"reflect"
	"runtime"
	"unsafe"

	"github.com/gowebapi/webapi/core/js"
)

{{end}}

{{define "type-uint8"}}

// {{.Name}}ToJs is converting []{{.GoType}} to a new javascript {{.Js}}.
func {{.Name}}ToJs(src []{{.GoType}} ) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * {{.Size}})
	js.CopyBytesToJS(array, src)
	return array
}


// JsTo{{.Name}} convert javascript {{.Js}} to Go {{.GoType}}.
func JsTo{{.Name}}(array js.Value) []{{.GoType}} {
	size := array.Get("length").Int()

	mem := make([]byte, size)
	js.CopyBytesToGo(mem, array)
	
	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	dst := *(*[]{{.GoType}})(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

{{end}}

{{define "type"}}

// {{.Name}}ToJs is converting []{{.GoType}} to a new javascript {{.Js}}.
func {{.Name}}ToJs(src []{{.GoType}} ) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * {{.Size}})

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
{{if gt .Size 1}}	head.Len *= {{.Size}}
	head.Cap *= {{.Size}}
{{end}}		data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("{{.Js}}").New(
		buf, 
		array.Get("byteOffset"), 
		array.Get("byteLength").Int()/{{.Size}},
	)
}

// JsTo{{.Name}} convert javascript {{.Js}} to Go {{.GoType}}.
func JsTo{{.Name}}(array js.Value) []{{.GoType}} {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size * {{.Size}})
	js.CopyBytesToGo(mem, tmp)
	
	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= {{.Size}}
	head.Len /= {{.Size}}
	dst := *(*[]{{.GoType}})(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
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
	TypeInfo{"uint8", "UInt8", "Uint8Array", 1},
	TypeInfo{"int8", "Int8", "Int8Array", 1},
	TypeInfo{"uint16", "UInt16", "Uint16Array", 2},
	TypeInfo{"int16", "Int16", "Int16Array", 2},
	TypeInfo{"uint32", "UInt32", "Uint32Array", 4},
	TypeInfo{"int32", "Int32", "Int32Array", 4},
	TypeInfo{"float32", "Float32", "Float32Array", 4},
	TypeInfo{"float64", "Float64", "Float64Array", 8},
}

/*
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
*/

func main() {
	content, err := generateText()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	if source, err := format.Source(content); err == nil {
		content = source
	} else {
		// we just print this error to get an output file that we
		// later can correct and fix the bug
		fmt.Fprintf(os.Stderr, "error:unable to format output source code: %s\n", err)
	}
	if err := ioutil.WriteFile("jsconv.go", content, 0664); err != nil {
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
