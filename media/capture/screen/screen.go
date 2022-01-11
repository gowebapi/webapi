// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package screen

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// screen-capture.idl

// transform files:
// screen-capture.go.md

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

// dictionary: DisplayMediaStreamConstraints
type DisplayMediaStreamConstraints struct {
	Video *Union
	Audio *Union
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *DisplayMediaStreamConstraints) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Video.JSValue()
	out.Set("video", value0)
	value1 := _this.Audio.JSValue()
	out.Set("audio", value1)
	return out
}

// DisplayMediaStreamConstraintsFromJS is allocating a new
// DisplayMediaStreamConstraints object and copy all values in the value javascript object.
func DisplayMediaStreamConstraintsFromJS(value js.Value) *DisplayMediaStreamConstraints {
	var out DisplayMediaStreamConstraints
	var (
		value0 *Union // javascript: Union {video Video video}
		value1 *Union // javascript: Union {audio Audio audio}
	)
	value0 = UnionFromJS(value.Get("video"))
	out.Video = value0
	value1 = UnionFromJS(value.Get("audio"))
	out.Audio = value1
	return &out
}
