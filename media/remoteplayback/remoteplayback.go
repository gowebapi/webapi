// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package remoteplayback

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.PromiseInt
// javascript.PromiseVoid

// source idl files:
// remote-playback.idl

// transform files:
// remote-playback.go.md

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

// enum: RemotePlaybackState
type RemotePlaybackState int

const (
	ConnectingRemotePlaybackState RemotePlaybackState = iota
	ConnectedRemotePlaybackState
	DisconnectedRemotePlaybackState
)

var remotePlaybackStateToWasmTable = []string{
	"connecting", "connected", "disconnected",
}

var remotePlaybackStateFromWasmTable = map[string]RemotePlaybackState{
	"connecting": ConnectingRemotePlaybackState, "connected": ConnectedRemotePlaybackState, "disconnected": DisconnectedRemotePlaybackState,
}

// JSValue is converting this enum into a javascript object
func (this *RemotePlaybackState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RemotePlaybackState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(remotePlaybackStateToWasmTable) {
		return remotePlaybackStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// RemotePlaybackStateFromJS is converting a javascript value into
// a RemotePlaybackState enum value.
func RemotePlaybackStateFromJS(value js.Value) RemotePlaybackState {
	key := value.String()
	conv, ok := remotePlaybackStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: RemotePlaybackAvailabilityCallback
type RemotePlaybackAvailabilityCallbackFunc func(available bool)

// RemotePlaybackAvailabilityCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type RemotePlaybackAvailabilityCallback js.Func

func RemotePlaybackAvailabilityCallbackToJS(callback RemotePlaybackAvailabilityCallbackFunc) *RemotePlaybackAvailabilityCallback {
	if callback == nil {
		return nil
	}
	ret := RemotePlaybackAvailabilityCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 bool // javascript: boolean available
		)
		_p0 = (args[0]).Bool()
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func RemotePlaybackAvailabilityCallbackFromJS(_value js.Value) RemotePlaybackAvailabilityCallbackFunc {
	return func(available bool) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := available
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// class: RemotePlayback
type RemotePlayback struct {
	domcore.EventTarget
}

// RemotePlaybackFromJS is casting a js.Wrapper into RemotePlayback.
func RemotePlaybackFromJS(value js.Wrapper) *RemotePlayback {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &RemotePlayback{}
	ret.Value_JS = input
	return ret
}

// State returning attribute 'state' with
// type RemotePlaybackState (idl: RemotePlaybackState).
func (_this *RemotePlayback) State() RemotePlaybackState {
	var ret RemotePlaybackState
	value := _this.Value_JS.Get("state")
	ret = RemotePlaybackStateFromJS(value)
	return ret
}

// OnConnecting returning attribute 'onconnecting' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) OnConnecting() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnecting")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnConnect returning attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) OnConnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnect")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnDisconnect returning attribute 'ondisconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) OnDisconnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ondisconnect")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: domcore.Event
func eventFuncRemotePlayback_domcore_Event(listener func(event *domcore.Event, target *RemotePlayback)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *domcore.Event
		value := args[0]
		incoming := value.Get("target")
		ret = domcore.EventFromJS(value)
		src := RemotePlaybackFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddConnect is adding doing AddEventListener for 'Connect' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) AddEventConnect(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "connect", cb)
	return cb
}

// SetOnConnect is assigning a function to 'onconnect'. This
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) SetOnConnect(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Set("onconnect", cb)
	return cb
}

// AddConnecting is adding doing AddEventListener for 'Connecting' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) AddEventConnecting(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "connecting", cb)
	return cb
}

// SetOnConnecting is assigning a function to 'onconnecting'. This
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) SetOnConnecting(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Set("onconnecting", cb)
	return cb
}

// AddDisconnect is adding doing AddEventListener for 'Disconnect' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) AddEventDisconnect(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "disconnect", cb)
	return cb
}

// SetOnDisconnect is assigning a function to 'ondisconnect'. This
// This method is returning allocated javascript function that need to be released.
func (_this *RemotePlayback) SetOnDisconnect(listener func(event *domcore.Event, currentTarget *RemotePlayback)) js.Func {
	cb := eventFuncRemotePlayback_domcore_Event(listener)
	_this.Value_JS.Set("ondisconnect", cb)
	return cb
}

func (_this *RemotePlayback) WatchAvailability(callback *RemotePlaybackAvailabilityCallback) (_result *javascript.PromiseInt) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("watchAvailability", _args[0:_end]...)
	var (
		_converted *javascript.PromiseInt // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseIntFromJS(_returned)
	_result = _converted
	return
}

func (_this *RemotePlayback) CancelWatchAvailability(id *int) (_result *javascript.PromiseVoid) {
	var (
		_args [1]interface{}
		_end  int
	)
	if id != nil {

		var _p0 interface{}
		if id != nil {
			_p0 = *(id)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("cancelWatchAvailability", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *RemotePlayback) Prompt() (_result *javascript.PromiseVoid) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("prompt", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}
