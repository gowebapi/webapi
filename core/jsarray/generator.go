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

package jsarray

//go:generate go run generator.go

import (
	"github.com/gowebapi/webapi/core/js"
)

{{end}}

{{define "type"}}

func {{.Name}}ToJS(values [] {{.GoType}} ) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func {{.Name}}ToGo(array js.Value) [] {{.GoType}} {
	size := array.Length()
	out := make( [] {{.GoType}} , size)
	for i := 0; i < size; i++ {
		out[i] = {{.GoType}} ( array.Index(i).{{.Method}} () )
	}
	return out
}

{{end}}
`

var tmpl = template.Must(template.New("main").Parse(templateText))

type TypeInfo struct {
	GoType string
	Name   string
	Method string
}

var types = []TypeInfo{
	TypeInfo{"uint8", "UInt8", "Int"},
	TypeInfo{"int8", "Int8", "Int"},
	TypeInfo{"uint16", "UInt16", "Int"},
	TypeInfo{"int16", "Int16", "Int"},
	TypeInfo{"uint32", "UInt32", "Int"},
	TypeInfo{"int32", "Int32", "Int"},
	TypeInfo{"float32", "Float32", "Float"},
	TypeInfo{"float64", "Float64", "Float"},
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
	if err := ioutil.WriteFile("jsarray.go", content, 0664); err != nil {
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
