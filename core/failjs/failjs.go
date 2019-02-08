// Equal api to syscall/js but is panic on all method invocation.
//
// The usage is to get a tab complession to work inside an IDE
// (e.g. Visual Studio Code) without changing to GOOS to js
package failjs

const message = "this library should be executed under WASM"

type Callback struct {
	Value
}

func NewCallback(fn func(args []Value)) Callback {
	panic(message)
}

func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) Callback {
	panic(message)
}

func (c Callback) Release() {
	panic(message)
}

type Error struct {
	Value
}

func (e Error) Error() string {
	panic(message)
}

type EventCallbackFlag int

const (
	PreventDefault EventCallbackFlag = 1 << iota
	StopPropagation
	StopImmediatePropagation
)

type Func struct {
	Value
}

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	panic(message)
}

func (c Func) Release() {
	panic(message)
}

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

type TypedArray struct {
	Value
}

func TypedArrayOf(slice interface{}) TypedArray {
	panic(message)
}

func (a TypedArray) Release() {
	panic(message)
}

type Value struct {
}

func Global() Value {
	panic(message)
}

func Null() Value {
	panic(message)
}

func Undefined() Value {
	panic(message)
}

func ValueOf(x interface{}) Value {
	panic(message)
}

func (v Value) Bool() bool {
	panic(message)
}

func (v Value) Call(m string, args ...interface{}) Value {
	panic(message)
}

func (v Value) Float() float64 {
	panic(message)
}

func (v Value) Get(p string) Value {
	panic(message)
}

func (v Value) Index(i int) Value {
	panic(message)
}

func (v Value) InstanceOf(t Value) bool {
	panic(message)
}

func (v Value) Int() int {
	panic(message)
}

func (v Value) Invoke(args ...interface{}) Value {
	panic(message)
}

func (v Value) JSValue() Value {
	panic(message)
}

func (v Value) Length() int {
	panic(message)
}

func (v Value) New(args ...interface{}) Value {
	panic(message)
}

func (v Value) Set(p string, x interface{}) {
	panic(message)
}

func (v Value) SetIndex(i int, x interface{}) {
	panic(message)
}

func (v Value) String() string {
	panic(message)
}

func (v Value) Truthy() bool {
	panic(message)
}

func (v Value) Type() Type {
	panic(message)
}

type ValueError struct {
	Method string
	Type   Type
}

func (e *ValueError) Error() string {
	panic(message)
}

type Wrapper interface {
	JSValue() Value
}
