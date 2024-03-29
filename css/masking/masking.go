// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package masking

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/graphics/svg"
)

// using following types:
// svg.SVGAnimatedEnumeration
// svg.SVGAnimatedLength
// svg.SVGAnimatedTransformList
// svg.SVGElement

// source idl files:
// css-masking.idl

// transform files:
// css-masking.go.md

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// class: SVGClipPathElement
type SVGClipPathElement struct {
	svg.SVGElement
}

// SVGClipPathElementFromJS is casting a js.Value into SVGClipPathElement.
func SVGClipPathElementFromJS(value js.Value) *SVGClipPathElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &SVGClipPathElement{}
	ret.Value_JS = value
	return ret
}

// SVGClipPathElementFromJS is casting from something that holds a js.Value into SVGClipPathElement.
func SVGClipPathElementFromWrapper(input core.Wrapper) *SVGClipPathElement {
	return SVGClipPathElementFromJS(input.JSValue())
}

// ClipPathUnits returning attribute 'clipPathUnits' with
// type svg.SVGAnimatedEnumeration (idl: SVGAnimatedEnumeration).
func (_this *SVGClipPathElement) ClipPathUnits() *svg.SVGAnimatedEnumeration {
	var ret *svg.SVGAnimatedEnumeration
	value := _this.Value_JS.Get("clipPathUnits")
	ret = svg.SVGAnimatedEnumerationFromJS(value)
	return ret
}

// Transform returning attribute 'transform' with
// type svg.SVGAnimatedTransformList (idl: SVGAnimatedTransformList).
func (_this *SVGClipPathElement) Transform() *svg.SVGAnimatedTransformList {
	var ret *svg.SVGAnimatedTransformList
	value := _this.Value_JS.Get("transform")
	ret = svg.SVGAnimatedTransformListFromJS(value)
	return ret
}

// class: SVGMaskElement
type SVGMaskElement struct {
	svg.SVGElement
}

// SVGMaskElementFromJS is casting a js.Value into SVGMaskElement.
func SVGMaskElementFromJS(value js.Value) *SVGMaskElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &SVGMaskElement{}
	ret.Value_JS = value
	return ret
}

// SVGMaskElementFromJS is casting from something that holds a js.Value into SVGMaskElement.
func SVGMaskElementFromWrapper(input core.Wrapper) *SVGMaskElement {
	return SVGMaskElementFromJS(input.JSValue())
}

// MaskUnits returning attribute 'maskUnits' with
// type svg.SVGAnimatedEnumeration (idl: SVGAnimatedEnumeration).
func (_this *SVGMaskElement) MaskUnits() *svg.SVGAnimatedEnumeration {
	var ret *svg.SVGAnimatedEnumeration
	value := _this.Value_JS.Get("maskUnits")
	ret = svg.SVGAnimatedEnumerationFromJS(value)
	return ret
}

// MaskContentUnits returning attribute 'maskContentUnits' with
// type svg.SVGAnimatedEnumeration (idl: SVGAnimatedEnumeration).
func (_this *SVGMaskElement) MaskContentUnits() *svg.SVGAnimatedEnumeration {
	var ret *svg.SVGAnimatedEnumeration
	value := _this.Value_JS.Get("maskContentUnits")
	ret = svg.SVGAnimatedEnumerationFromJS(value)
	return ret
}

// X returning attribute 'x' with
// type svg.SVGAnimatedLength (idl: SVGAnimatedLength).
func (_this *SVGMaskElement) X() *svg.SVGAnimatedLength {
	var ret *svg.SVGAnimatedLength
	value := _this.Value_JS.Get("x")
	ret = svg.SVGAnimatedLengthFromJS(value)
	return ret
}

// Y returning attribute 'y' with
// type svg.SVGAnimatedLength (idl: SVGAnimatedLength).
func (_this *SVGMaskElement) Y() *svg.SVGAnimatedLength {
	var ret *svg.SVGAnimatedLength
	value := _this.Value_JS.Get("y")
	ret = svg.SVGAnimatedLengthFromJS(value)
	return ret
}

// Width returning attribute 'width' with
// type svg.SVGAnimatedLength (idl: SVGAnimatedLength).
func (_this *SVGMaskElement) Width() *svg.SVGAnimatedLength {
	var ret *svg.SVGAnimatedLength
	value := _this.Value_JS.Get("width")
	ret = svg.SVGAnimatedLengthFromJS(value)
	return ret
}

// Height returning attribute 'height' with
// type svg.SVGAnimatedLength (idl: SVGAnimatedLength).
func (_this *SVGMaskElement) Height() *svg.SVGAnimatedLength {
	var ret *svg.SVGAnimatedLength
	value := _this.Value_JS.Get("height")
	ret = svg.SVGAnimatedLengthFromJS(value)
	return ret
}
