// Code generated by webidl-bind. DO NOT EDIT.

package intersection

import "syscall/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/geometry"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// dom.Element
// geometry.DOMRectInit
// geometry.DOMRectReadOnly
// javascript.FrozenArray

// source idl files:
// intersection-observer.idl

// transform files:
// intersection-observer.go.md

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

// callback: IntersectionObserverCallback
type IntersectionObserverCallbackFunc func(entries []*IntersectionObserverEntry, observer *IntersectionObserver)

// IntersectionObserverCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type IntersectionObserverCallback js.Func

func IntersectionObserverCallbackToJS(callback IntersectionObserverCallbackFunc) *IntersectionObserverCallback {
	if callback == nil {
		return nil
	}
	ret := IntersectionObserverCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*IntersectionObserverEntry // javascript: sequence<IntersectionObserverEntry> entries
			_p1 *IntersectionObserver        // javascript: IntersectionObserver observer
		)
		__length0 := args[0].Length()
		__array0 := make([]*IntersectionObserverEntry, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *IntersectionObserverEntry
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = IntersectionObserverEntryFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		_p1 = IntersectionObserverFromJS(args[1])
		callback(_p0, _p1)

		// returning no return value
		return nil
	}))
	return &ret
}

func IntersectionObserverCallbackFromJS(_value js.Value) IntersectionObserverCallbackFunc {
	return func(entries []*IntersectionObserverEntry, observer *IntersectionObserver) {
		var (
			_args [2]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(entries))
		for __idx0, __seq_in0 := range entries {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_p1 := observer.JSValue()
		_args[1] = _p1
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: IntersectionObserverEntryInit
type IntersectionObserverEntryInit struct {
	Time               float64
	RootBounds         *geometry.DOMRectInit
	BoundingClientRect *geometry.DOMRectInit
	IntersectionRect   *geometry.DOMRectInit
	IsIntersecting     bool
	IntersectionRatio  float64
	Target             *dom.Element
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *IntersectionObserverEntryInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Time
	out.Set("time", value0)
	value1 := _this.RootBounds.JSValue()
	out.Set("rootBounds", value1)
	value2 := _this.BoundingClientRect.JSValue()
	out.Set("boundingClientRect", value2)
	value3 := _this.IntersectionRect.JSValue()
	out.Set("intersectionRect", value3)
	value4 := _this.IsIntersecting
	out.Set("isIntersecting", value4)
	value5 := _this.IntersectionRatio
	out.Set("intersectionRatio", value5)
	value6 := _this.Target.JSValue()
	out.Set("target", value6)
	return out
}

// IntersectionObserverEntryInitFromJS is allocating a new
// IntersectionObserverEntryInit object and copy all values in the value javascript object.
func IntersectionObserverEntryInitFromJS(value js.Value) *IntersectionObserverEntryInit {
	var out IntersectionObserverEntryInit
	var (
		value0 float64               // javascript: double {time Time time}
		value1 *geometry.DOMRectInit // javascript: DOMRectInit {rootBounds RootBounds rootBounds}
		value2 *geometry.DOMRectInit // javascript: DOMRectInit {boundingClientRect BoundingClientRect boundingClientRect}
		value3 *geometry.DOMRectInit // javascript: DOMRectInit {intersectionRect IntersectionRect intersectionRect}
		value4 bool                  // javascript: boolean {isIntersecting IsIntersecting isIntersecting}
		value5 float64               // javascript: double {intersectionRatio IntersectionRatio intersectionRatio}
		value6 *dom.Element          // javascript: Element {target Target target}
	)
	value0 = (value.Get("time")).Float()
	out.Time = value0
	if value.Get("rootBounds").Type() != js.TypeNull && value.Get("rootBounds").Type() != js.TypeUndefined {
		value1 = geometry.DOMRectInitFromJS(value.Get("rootBounds"))
	}
	out.RootBounds = value1
	value2 = geometry.DOMRectInitFromJS(value.Get("boundingClientRect"))
	out.BoundingClientRect = value2
	value3 = geometry.DOMRectInitFromJS(value.Get("intersectionRect"))
	out.IntersectionRect = value3
	value4 = (value.Get("isIntersecting")).Bool()
	out.IsIntersecting = value4
	value5 = (value.Get("intersectionRatio")).Float()
	out.IntersectionRatio = value5
	value6 = dom.ElementFromJS(value.Get("target"))
	out.Target = value6
	return &out
}

// dictionary: IntersectionObserverInit
type IntersectionObserverInit struct {
	Root       *dom.Element
	RootMargin string
	Threshold  *Union
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *IntersectionObserverInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Root.JSValue()
	out.Set("root", value0)
	value1 := _this.RootMargin
	out.Set("rootMargin", value1)
	value2 := _this.Threshold.JSValue()
	out.Set("threshold", value2)
	return out
}

// IntersectionObserverInitFromJS is allocating a new
// IntersectionObserverInit object and copy all values in the value javascript object.
func IntersectionObserverInitFromJS(value js.Value) *IntersectionObserverInit {
	var out IntersectionObserverInit
	var (
		value0 *dom.Element // javascript: Element {root Root root}
		value1 string       // javascript: DOMString {rootMargin RootMargin rootMargin}
		value2 *Union       // javascript: Union {threshold Threshold threshold}
	)
	if value.Get("root").Type() != js.TypeNull && value.Get("root").Type() != js.TypeUndefined {
		value0 = dom.ElementFromJS(value.Get("root"))
	}
	out.Root = value0
	value1 = (value.Get("rootMargin")).String()
	out.RootMargin = value1
	value2 = UnionFromJS(value.Get("threshold"))
	out.Threshold = value2
	return &out
}

// class: IntersectionObserver
type IntersectionObserver struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *IntersectionObserver) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// IntersectionObserverFromJS is casting a js.Value into IntersectionObserver.
func IntersectionObserverFromJS(value js.Value) *IntersectionObserver {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &IntersectionObserver{}
	ret.Value_JS = value
	return ret
}

// IntersectionObserverFromJS is casting from something that holds a js.Value into IntersectionObserver.
func IntersectionObserverFromWrapper(input core.Wrapper) *IntersectionObserver {
	return IntersectionObserverFromJS(input.JSValue())
}

func NewIntersectionObserver(callback *IntersectionObserverCallback, options *IntersectionObserverInit) (_result *IntersectionObserver) {
	_klass := js.Global().Get("IntersectionObserver")
	var (
		_args [2]interface{}
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
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *IntersectionObserver // javascript: IntersectionObserver _what_return_name
	)
	_converted = IntersectionObserverFromJS(_returned)
	_result = _converted
	return
}

// Root returning attribute 'root' with
// type dom.Element (idl: Element).
func (_this *IntersectionObserver) Root() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("root")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = dom.ElementFromJS(value)
	}
	return ret
}

// RootMargin returning attribute 'rootMargin' with
// type string (idl: DOMString).
func (_this *IntersectionObserver) RootMargin() string {
	var ret string
	value := _this.Value_JS.Get("rootMargin")
	ret = (value).String()
	return ret
}

// Thresholds returning attribute 'thresholds' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *IntersectionObserver) Thresholds() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("thresholds")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *IntersectionObserver) Observe(target *dom.Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("observe", _args[0:_end]...)
	return
}

func (_this *IntersectionObserver) Unobserve(target *dom.Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("unobserve", _args[0:_end]...)
	return
}

func (_this *IntersectionObserver) Disconnect() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("disconnect", _args[0:_end]...)
	return
}

func (_this *IntersectionObserver) TakeRecords() (_result []*IntersectionObserverEntry) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("takeRecords", _args[0:_end]...)
	var (
		_converted []*IntersectionObserverEntry // javascript: sequence<IntersectionObserverEntry> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*IntersectionObserverEntry, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *IntersectionObserverEntry
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = IntersectionObserverEntryFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

// class: IntersectionObserverEntry
type IntersectionObserverEntry struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *IntersectionObserverEntry) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// IntersectionObserverEntryFromJS is casting a js.Value into IntersectionObserverEntry.
func IntersectionObserverEntryFromJS(value js.Value) *IntersectionObserverEntry {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &IntersectionObserverEntry{}
	ret.Value_JS = value
	return ret
}

// IntersectionObserverEntryFromJS is casting from something that holds a js.Value into IntersectionObserverEntry.
func IntersectionObserverEntryFromWrapper(input core.Wrapper) *IntersectionObserverEntry {
	return IntersectionObserverEntryFromJS(input.JSValue())
}

func NewIntersectionObserverEntry(intersectionObserverEntryInit *IntersectionObserverEntryInit) (_result *IntersectionObserverEntry) {
	_klass := js.Global().Get("IntersectionObserverEntry")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := intersectionObserverEntryInit.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *IntersectionObserverEntry // javascript: IntersectionObserverEntry _what_return_name
	)
	_converted = IntersectionObserverEntryFromJS(_returned)
	_result = _converted
	return
}

// Time returning attribute 'time' with
// type float64 (idl: double).
func (_this *IntersectionObserverEntry) Time() float64 {
	var ret float64
	value := _this.Value_JS.Get("time")
	ret = (value).Float()
	return ret
}

// RootBounds returning attribute 'rootBounds' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *IntersectionObserverEntry) RootBounds() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("rootBounds")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = geometry.DOMRectReadOnlyFromJS(value)
	}
	return ret
}

// BoundingClientRect returning attribute 'boundingClientRect' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *IntersectionObserverEntry) BoundingClientRect() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("boundingClientRect")
	ret = geometry.DOMRectReadOnlyFromJS(value)
	return ret
}

// IntersectionRect returning attribute 'intersectionRect' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *IntersectionObserverEntry) IntersectionRect() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("intersectionRect")
	ret = geometry.DOMRectReadOnlyFromJS(value)
	return ret
}

// IsIntersecting returning attribute 'isIntersecting' with
// type bool (idl: boolean).
func (_this *IntersectionObserverEntry) IsIntersecting() bool {
	var ret bool
	value := _this.Value_JS.Get("isIntersecting")
	ret = (value).Bool()
	return ret
}

// IntersectionRatio returning attribute 'intersectionRatio' with
// type float64 (idl: double).
func (_this *IntersectionObserverEntry) IntersectionRatio() float64 {
	var ret float64
	value := _this.Value_JS.Get("intersectionRatio")
	ret = (value).Float()
	return ret
}

// Target returning attribute 'target' with
// type dom.Element (idl: Element).
func (_this *IntersectionObserverEntry) Target() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("target")
	ret = dom.ElementFromJS(value)
	return ret
}
