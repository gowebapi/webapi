// Code generated by webidl-bind. DO NOT EDIT.

package view

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/dom/geometry"
	"github.com/gowebapi/webapi/media/capabilities"
	"github.com/gowebapi/webapi/media/orientation"
)

// using following types:
// capabilities.ScreenColorGamut
// capabilities.ScreenLuminance
// domcore.Event
// domcore.EventHandler
// domcore.EventListener
// domcore.EventTarget
// geometry.DOMRect
// orientation.ScreenOrientation

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

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

// enum: CSSBoxType
type CSSBoxType int

const (
	MarginCSSBoxType CSSBoxType = iota
	BorderCSSBoxType
	PaddingCSSBoxType
	ContentCSSBoxType
)

var cSSBoxTypeToWasmTable = []string{
	"margin", "border", "padding", "content",
}

var cSSBoxTypeFromWasmTable = map[string]CSSBoxType{
	"margin": MarginCSSBoxType, "border": BorderCSSBoxType, "padding": PaddingCSSBoxType, "content": ContentCSSBoxType,
}

// JSValue is converting this enum into a java object
func (this *CSSBoxType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this CSSBoxType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(cSSBoxTypeToWasmTable) {
		return cSSBoxTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// CSSBoxTypeFromJS is converting a javascript value into
// a CSSBoxType enum value.
func CSSBoxTypeFromJS(value js.Value) CSSBoxType {
	key := value.String()
	conv, ok := cSSBoxTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ScrollBehavior
type ScrollBehavior int

const (
	AutoScrollBehavior ScrollBehavior = iota
	SmoothScrollBehavior
)

var scrollBehaviorToWasmTable = []string{
	"auto", "smooth",
}

var scrollBehaviorFromWasmTable = map[string]ScrollBehavior{
	"auto": AutoScrollBehavior, "smooth": SmoothScrollBehavior,
}

// JSValue is converting this enum into a java object
func (this *ScrollBehavior) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollBehavior) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollBehaviorToWasmTable) {
		return scrollBehaviorToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollBehaviorFromJS is converting a javascript value into
// a ScrollBehavior enum value.
func ScrollBehaviorFromJS(value js.Value) ScrollBehavior {
	key := value.String()
	conv, ok := scrollBehaviorFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ScrollLogicalPosition
type ScrollLogicalPosition int

const (
	StartScrollLogicalPosition ScrollLogicalPosition = iota
	CenterScrollLogicalPosition
	EndScrollLogicalPosition
	NearestScrollLogicalPosition
)

var scrollLogicalPositionToWasmTable = []string{
	"start", "center", "end", "nearest",
}

var scrollLogicalPositionFromWasmTable = map[string]ScrollLogicalPosition{
	"start": StartScrollLogicalPosition, "center": CenterScrollLogicalPosition, "end": EndScrollLogicalPosition, "nearest": NearestScrollLogicalPosition,
}

// JSValue is converting this enum into a java object
func (this *ScrollLogicalPosition) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollLogicalPosition) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollLogicalPositionToWasmTable) {
		return scrollLogicalPositionToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollLogicalPositionFromJS is converting a javascript value into
// a ScrollLogicalPosition enum value.
func ScrollLogicalPositionFromJS(value js.Value) ScrollLogicalPosition {
	key := value.String()
	conv, ok := scrollLogicalPositionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: BoxQuadOptions
type BoxQuadOptions struct {
	Box        CSSBoxType
	RelativeTo *Union
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BoxQuadOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Box.JSValue()
	out.Set("box", value0)
	value1 := _this.RelativeTo.JSValue()
	out.Set("relativeTo", value1)
	return out
}

// BoxQuadOptionsFromJS is allocating a new
// BoxQuadOptions object and copy all values from
// input javascript object
func BoxQuadOptionsFromJS(value js.Wrapper) *BoxQuadOptions {
	input := value.JSValue()
	var out BoxQuadOptions
	var (
		value0 CSSBoxType // javascript: CSSBoxType {box Box box}
		value1 *Union     // javascript: Union {relativeTo RelativeTo relativeTo}
	)
	value0 = CSSBoxTypeFromJS(input.Get("box"))
	out.Box = value0
	value1 = UnionFromJS(input.Get("relativeTo"))
	out.RelativeTo = value1
	return &out
}

// dictionary: ConvertCoordinateOptions
type ConvertCoordinateOptions struct {
	FromBox CSSBoxType
	ToBox   CSSBoxType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ConvertCoordinateOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.FromBox.JSValue()
	out.Set("fromBox", value0)
	value1 := _this.ToBox.JSValue()
	out.Set("toBox", value1)
	return out
}

// ConvertCoordinateOptionsFromJS is allocating a new
// ConvertCoordinateOptions object and copy all values from
// input javascript object
func ConvertCoordinateOptionsFromJS(value js.Wrapper) *ConvertCoordinateOptions {
	input := value.JSValue()
	var out ConvertCoordinateOptions
	var (
		value0 CSSBoxType // javascript: CSSBoxType {fromBox FromBox fromBox}
		value1 CSSBoxType // javascript: CSSBoxType {toBox ToBox toBox}
	)
	value0 = CSSBoxTypeFromJS(input.Get("fromBox"))
	out.FromBox = value0
	value1 = CSSBoxTypeFromJS(input.Get("toBox"))
	out.ToBox = value1
	return &out
}

// dictionary: MediaQueryListEventInit
type MediaQueryListEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Media      string
	Matches    bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaQueryListEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Media
	out.Set("media", value3)
	value4 := _this.Matches
	out.Set("matches", value4)
	return out
}

// MediaQueryListEventInitFromJS is allocating a new
// MediaQueryListEventInit object and copy all values from
// input javascript object
func MediaQueryListEventInitFromJS(value js.Wrapper) *MediaQueryListEventInit {
	input := value.JSValue()
	var out MediaQueryListEventInit
	var (
		value0 bool   // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool   // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool   // javascript: boolean {composed Composed composed}
		value3 string // javascript: DOMString {media Media media}
		value4 bool   // javascript: boolean {matches Matches matches}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("media")).String()
	out.Media = value3
	value4 = (input.Get("matches")).Bool()
	out.Matches = value4
	return &out
}

// dictionary: ScrollIntoViewOptions
type ScrollIntoViewOptions struct {
	Behavior ScrollBehavior
	Block    ScrollLogicalPosition
	Inline   ScrollLogicalPosition
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ScrollIntoViewOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Behavior.JSValue()
	out.Set("behavior", value0)
	value1 := _this.Block.JSValue()
	out.Set("block", value1)
	value2 := _this.Inline.JSValue()
	out.Set("inline", value2)
	return out
}

// ScrollIntoViewOptionsFromJS is allocating a new
// ScrollIntoViewOptions object and copy all values from
// input javascript object
func ScrollIntoViewOptionsFromJS(value js.Wrapper) *ScrollIntoViewOptions {
	input := value.JSValue()
	var out ScrollIntoViewOptions
	var (
		value0 ScrollBehavior        // javascript: ScrollBehavior {behavior Behavior behavior}
		value1 ScrollLogicalPosition // javascript: ScrollLogicalPosition {block Block block}
		value2 ScrollLogicalPosition // javascript: ScrollLogicalPosition {inline Inline inline}
	)
	value0 = ScrollBehaviorFromJS(input.Get("behavior"))
	out.Behavior = value0
	value1 = ScrollLogicalPositionFromJS(input.Get("block"))
	out.Block = value1
	value2 = ScrollLogicalPositionFromJS(input.Get("inline"))
	out.Inline = value2
	return &out
}

// dictionary: ScrollOptions
type ScrollOptions struct {
	Behavior ScrollBehavior
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ScrollOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Behavior.JSValue()
	out.Set("behavior", value0)
	return out
}

// ScrollOptionsFromJS is allocating a new
// ScrollOptions object and copy all values from
// input javascript object
func ScrollOptionsFromJS(value js.Wrapper) *ScrollOptions {
	input := value.JSValue()
	var out ScrollOptions
	var (
		value0 ScrollBehavior // javascript: ScrollBehavior {behavior Behavior behavior}
	)
	value0 = ScrollBehaviorFromJS(input.Get("behavior"))
	out.Behavior = value0
	return &out
}

// dictionary: ScrollToOptions
type ScrollToOptions struct {
	Behavior ScrollBehavior
	Left     float64
	Top      float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ScrollToOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Behavior.JSValue()
	out.Set("behavior", value0)
	value1 := _this.Left
	out.Set("left", value1)
	value2 := _this.Top
	out.Set("top", value2)
	return out
}

// ScrollToOptionsFromJS is allocating a new
// ScrollToOptions object and copy all values from
// input javascript object
func ScrollToOptionsFromJS(value js.Wrapper) *ScrollToOptions {
	input := value.JSValue()
	var out ScrollToOptions
	var (
		value0 ScrollBehavior // javascript: ScrollBehavior {behavior Behavior behavior}
		value1 float64        // javascript: unrestricted double {left Left left}
		value2 float64        // javascript: unrestricted double {top Top top}
	)
	value0 = ScrollBehaviorFromJS(input.Get("behavior"))
	out.Behavior = value0
	value1 = (input.Get("left")).Float()
	out.Left = value1
	value2 = (input.Get("top")).Float()
	out.Top = value2
	return &out
}

// interface: CaretPosition
type CaretPosition struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CaretPosition) JSValue() js.Value {
	return _this.Value_JS
}

// CaretPositionFromJS is casting a js.Wrapper into CaretPosition.
func CaretPositionFromJS(value js.Wrapper) *CaretPosition {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CaretPosition{}
	ret.Value_JS = input
	return ret
}

// OffsetNode returning attribute 'offsetNode' with
// type js.Value (idl: <rawjs>).
func (_this *CaretPosition) OffsetNode() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("offsetNode")
	ret = value
	return ret
}

// Offset returning attribute 'offset' with
// type uint (idl: unsigned long).
func (_this *CaretPosition) Offset() uint {
	var ret uint
	value := _this.Value_JS.Get("offset")
	ret = (uint)((value).Int())
	return ret
}

func (_this *CaretPosition) GetClientRect() (_result *geometry.DOMRect) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getClientRect", _args[0:_end]...)
	var (
		_converted *geometry.DOMRect // javascript: DOMRect _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = geometry.DOMRectFromJS(_returned)
	}
	_result = _converted
	return
}

// interface: MediaQueryList
type MediaQueryList struct {
	domcore.EventTarget
}

// MediaQueryListFromJS is casting a js.Wrapper into MediaQueryList.
func MediaQueryListFromJS(value js.Wrapper) *MediaQueryList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaQueryList{}
	ret.Value_JS = input
	return ret
}

// Media returning attribute 'media' with
// type string (idl: DOMString).
func (_this *MediaQueryList) Media() string {
	var ret string
	value := _this.Value_JS.Get("media")
	ret = (value).String()
	return ret
}

// Matches returning attribute 'matches' with
// type bool (idl: boolean).
func (_this *MediaQueryList) Matches() bool {
	var ret bool
	value := _this.Value_JS.Get("matches")
	ret = (value).Bool()
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaQueryList) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaQueryList) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

func (_this *MediaQueryList) AddListener(listener *domcore.EventListenerValue) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := listener.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("addListener", _args[0:_end]...)
	return
}

func (_this *MediaQueryList) RemoveListener(listener *domcore.EventListenerValue) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := listener.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("removeListener", _args[0:_end]...)
	return
}

// interface: MediaQueryListEvent
type MediaQueryListEvent struct {
	domcore.Event
}

// MediaQueryListEventFromJS is casting a js.Wrapper into MediaQueryListEvent.
func MediaQueryListEventFromJS(value js.Wrapper) *MediaQueryListEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaQueryListEvent{}
	ret.Value_JS = input
	return ret
}

func NewMediaQueryListEvent(_type string, eventInitDict *MediaQueryListEventInit) (_result *MediaQueryListEvent) {
	_klass := js.Global().Get("MediaQueryListEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *MediaQueryListEvent // javascript: MediaQueryListEvent _what_return_name
	)
	_converted = MediaQueryListEventFromJS(_returned)
	_result = _converted
	return
}

// Media returning attribute 'media' with
// type string (idl: DOMString).
func (_this *MediaQueryListEvent) Media() string {
	var ret string
	value := _this.Value_JS.Get("media")
	ret = (value).String()
	return ret
}

// Matches returning attribute 'matches' with
// type bool (idl: boolean).
func (_this *MediaQueryListEvent) Matches() bool {
	var ret bool
	value := _this.Value_JS.Get("matches")
	ret = (value).Bool()
	return ret
}

// interface: Screen
type Screen struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Screen) JSValue() js.Value {
	return _this.Value_JS
}

// ScreenFromJS is casting a js.Wrapper into Screen.
func ScreenFromJS(value js.Wrapper) *Screen {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Screen{}
	ret.Value_JS = input
	return ret
}

// AvailWidth returning attribute 'availWidth' with
// type int (idl: long).
func (_this *Screen) AvailWidth() int {
	var ret int
	value := _this.Value_JS.Get("availWidth")
	ret = (value).Int()
	return ret
}

// AvailHeight returning attribute 'availHeight' with
// type int (idl: long).
func (_this *Screen) AvailHeight() int {
	var ret int
	value := _this.Value_JS.Get("availHeight")
	ret = (value).Int()
	return ret
}

// Width returning attribute 'width' with
// type int (idl: long).
func (_this *Screen) Width() int {
	var ret int
	value := _this.Value_JS.Get("width")
	ret = (value).Int()
	return ret
}

// Height returning attribute 'height' with
// type int (idl: long).
func (_this *Screen) Height() int {
	var ret int
	value := _this.Value_JS.Get("height")
	ret = (value).Int()
	return ret
}

// ColorDepth returning attribute 'colorDepth' with
// type uint (idl: unsigned long).
func (_this *Screen) ColorDepth() uint {
	var ret uint
	value := _this.Value_JS.Get("colorDepth")
	ret = (uint)((value).Int())
	return ret
}

// PixelDepth returning attribute 'pixelDepth' with
// type uint (idl: unsigned long).
func (_this *Screen) PixelDepth() uint {
	var ret uint
	value := _this.Value_JS.Get("pixelDepth")
	ret = (uint)((value).Int())
	return ret
}

// ColorGamut returning attribute 'colorGamut' with
// type capabilities.ScreenColorGamut (idl: ScreenColorGamut).
func (_this *Screen) ColorGamut() capabilities.ScreenColorGamut {
	var ret capabilities.ScreenColorGamut
	value := _this.Value_JS.Get("colorGamut")
	ret = capabilities.ScreenColorGamutFromJS(value)
	return ret
}

// Luminance returning attribute 'luminance' with
// type capabilities.ScreenLuminance (idl: ScreenLuminance).
func (_this *Screen) Luminance() *capabilities.ScreenLuminance {
	var ret *capabilities.ScreenLuminance
	value := _this.Value_JS.Get("luminance")
	if value.Type() != js.TypeNull {
		ret = capabilities.ScreenLuminanceFromJS(value)
	}
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Screen) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Screen) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

// Orientation returning attribute 'orientation' with
// type orientation.ScreenOrientation (idl: ScreenOrientation).
func (_this *Screen) Orientation() *orientation.ScreenOrientation {
	var ret *orientation.ScreenOrientation
	value := _this.Value_JS.Get("orientation")
	ret = orientation.ScreenOrientationFromJS(value)
	return ret
}
