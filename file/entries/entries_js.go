// Code generated by webidl-bind. DO NOT EDIT.

package entries

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/file"
)

// using following types:
// domcore.DOMException
// file.File

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

// callback: ErrorCallback
type ErrorCallbackFunc func(err *domcore.DOMException)

// ErrorCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type ErrorCallback js.Func

func ErrorCallbackToJS(callback ErrorCallbackFunc) *ErrorCallback {
	if callback == nil {
		return nil
	}
	ret := ErrorCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *domcore.DOMException // javascript: DOMException err
		)
		_p0 = domcore.DOMExceptionFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func ErrorCallbackFromJS(_value js.Value) ErrorCallbackFunc {
	return func(err *domcore.DOMException) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := err.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: FileCallback
type FileCallbackFunc func(file *file.File)

// FileCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FileCallback js.Func

func FileCallbackToJS(callback FileCallbackFunc) *FileCallback {
	if callback == nil {
		return nil
	}
	ret := FileCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *file.File // javascript: File file
		)
		_p0 = file.FileFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FileCallbackFromJS(_value js.Value) FileCallbackFunc {
	return func(file *file.File) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := file.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: FileSystemEntriesCallback
type FileSystemEntriesCallbackFunc func(entries []*FileSystemEntry)

// FileSystemEntriesCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FileSystemEntriesCallback js.Func

func FileSystemEntriesCallbackToJS(callback FileSystemEntriesCallbackFunc) *FileSystemEntriesCallback {
	if callback == nil {
		return nil
	}
	ret := FileSystemEntriesCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*FileSystemEntry // javascript: sequence<FileSystemEntry> entries
		)
		__length0 := args[0].Length()
		__array0 := make([]*FileSystemEntry, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *FileSystemEntry
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = FileSystemEntryFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FileSystemEntriesCallbackFromJS(_value js.Value) FileSystemEntriesCallbackFunc {
	return func(entries []*FileSystemEntry) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(entries))
		for __idx0, __seq_in0 := range entries {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: FileSystemEntryCallback
type FileSystemEntryCallbackFunc func(entry *FileSystemEntry)

// FileSystemEntryCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FileSystemEntryCallback js.Func

func FileSystemEntryCallbackToJS(callback FileSystemEntryCallbackFunc) *FileSystemEntryCallback {
	if callback == nil {
		return nil
	}
	ret := FileSystemEntryCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *FileSystemEntry // javascript: FileSystemEntry entry
		)
		_p0 = FileSystemEntryFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FileSystemEntryCallbackFromJS(_value js.Value) FileSystemEntryCallbackFunc {
	return func(entry *FileSystemEntry) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := entry.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: FileSystemFlags
type FileSystemFlags struct {
	Create    bool
	Exclusive bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FileSystemFlags) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Create
	out.Set("create", value0)
	value1 := _this.Exclusive
	out.Set("exclusive", value1)
	return out
}

// FileSystemFlagsFromJS is allocating a new
// FileSystemFlags object and copy all values from
// input javascript object
func FileSystemFlagsFromJS(value js.Wrapper) *FileSystemFlags {
	input := value.JSValue()
	var out FileSystemFlags
	var (
		value0 bool // javascript: boolean {create Create create}
		value1 bool // javascript: boolean {exclusive Exclusive exclusive}
	)
	value0 = (input.Get("create")).Bool()
	out.Create = value0
	value1 = (input.Get("exclusive")).Bool()
	out.Exclusive = value1
	return &out
}

// interface: FileSystem
type FileSystem struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileSystem) JSValue() js.Value {
	return _this.Value_JS
}

// FileSystemFromJS is casting a js.Wrapper into FileSystem.
func FileSystemFromJS(value js.Wrapper) *FileSystem {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileSystem{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: USVString).
func (_this *FileSystem) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Root returning attribute 'root' with
// type FileSystemDirectoryEntry (idl: FileSystemDirectoryEntry).
func (_this *FileSystem) Root() *FileSystemDirectoryEntry {
	var ret *FileSystemDirectoryEntry
	value := _this.Value_JS.Get("root")
	ret = FileSystemDirectoryEntryFromJS(value)
	return ret
}

// interface: FileSystemDirectoryEntry
type FileSystemDirectoryEntry struct {
	FileSystemEntry
}

// FileSystemDirectoryEntryFromJS is casting a js.Wrapper into FileSystemDirectoryEntry.
func FileSystemDirectoryEntryFromJS(value js.Wrapper) *FileSystemDirectoryEntry {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileSystemDirectoryEntry{}
	ret.Value_JS = input
	return ret
}

func (_this *FileSystemDirectoryEntry) CreateReader() (_result *FileSystemDirectoryReader) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("createReader", _args[0:_end]...)
	var (
		_converted *FileSystemDirectoryReader // javascript: FileSystemDirectoryReader _what_return_name
	)
	_converted = FileSystemDirectoryReaderFromJS(_returned)
	_result = _converted
	return
}

func (_this *FileSystemDirectoryEntry) GetFile(path *string, options *FileSystemFlags, successCallback *FileSystemEntryCallback, errorCallback *ErrorCallback) {
	var (
		_args [4]interface{}
		_end  int
	)
	if path != nil {
		_p0 := path
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	if successCallback != nil {

		var __callback2 js.Value
		if successCallback != nil {
			__callback2 = (*successCallback).Value
		} else {
			__callback2 = js.Null()
		}
		_p2 := __callback2
		_args[2] = _p2
		_end++
	}
	if errorCallback != nil {

		var __callback3 js.Value
		if errorCallback != nil {
			__callback3 = (*errorCallback).Value
		} else {
			__callback3 = js.Null()
		}
		_p3 := __callback3
		_args[3] = _p3
		_end++
	}
	_this.Value_JS.Call("getFile", _args[0:_end]...)
	return
}

func (_this *FileSystemDirectoryEntry) GetDirectory(path *string, options *FileSystemFlags, successCallback *FileSystemEntryCallback, errorCallback *ErrorCallback) {
	var (
		_args [4]interface{}
		_end  int
	)
	if path != nil {
		_p0 := path
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	if successCallback != nil {

		var __callback2 js.Value
		if successCallback != nil {
			__callback2 = (*successCallback).Value
		} else {
			__callback2 = js.Null()
		}
		_p2 := __callback2
		_args[2] = _p2
		_end++
	}
	if errorCallback != nil {

		var __callback3 js.Value
		if errorCallback != nil {
			__callback3 = (*errorCallback).Value
		} else {
			__callback3 = js.Null()
		}
		_p3 := __callback3
		_args[3] = _p3
		_end++
	}
	_this.Value_JS.Call("getDirectory", _args[0:_end]...)
	return
}

// interface: FileSystemDirectoryReader
type FileSystemDirectoryReader struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileSystemDirectoryReader) JSValue() js.Value {
	return _this.Value_JS
}

// FileSystemDirectoryReaderFromJS is casting a js.Wrapper into FileSystemDirectoryReader.
func FileSystemDirectoryReaderFromJS(value js.Wrapper) *FileSystemDirectoryReader {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileSystemDirectoryReader{}
	ret.Value_JS = input
	return ret
}

func (_this *FileSystemDirectoryReader) ReadEntries(successCallback *FileSystemEntriesCallback, errorCallback *ErrorCallback) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if successCallback != nil {
		__callback0 = (*successCallback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if errorCallback != nil {

		var __callback1 js.Value
		if errorCallback != nil {
			__callback1 = (*errorCallback).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("readEntries", _args[0:_end]...)
	return
}

// interface: FileSystemEntry
type FileSystemEntry struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileSystemEntry) JSValue() js.Value {
	return _this.Value_JS
}

// FileSystemEntryFromJS is casting a js.Wrapper into FileSystemEntry.
func FileSystemEntryFromJS(value js.Wrapper) *FileSystemEntry {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileSystemEntry{}
	ret.Value_JS = input
	return ret
}

// IsFile returning attribute 'isFile' with
// type bool (idl: boolean).
func (_this *FileSystemEntry) IsFile() bool {
	var ret bool
	value := _this.Value_JS.Get("isFile")
	ret = (value).Bool()
	return ret
}

// IsDirectory returning attribute 'isDirectory' with
// type bool (idl: boolean).
func (_this *FileSystemEntry) IsDirectory() bool {
	var ret bool
	value := _this.Value_JS.Get("isDirectory")
	ret = (value).Bool()
	return ret
}

// Name returning attribute 'name' with
// type string (idl: USVString).
func (_this *FileSystemEntry) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// FullPath returning attribute 'fullPath' with
// type string (idl: USVString).
func (_this *FileSystemEntry) FullPath() string {
	var ret string
	value := _this.Value_JS.Get("fullPath")
	ret = (value).String()
	return ret
}

// Filesystem returning attribute 'filesystem' with
// type FileSystem (idl: FileSystem).
func (_this *FileSystemEntry) Filesystem() *FileSystem {
	var ret *FileSystem
	value := _this.Value_JS.Get("filesystem")
	ret = FileSystemFromJS(value)
	return ret
}

func (_this *FileSystemEntry) GetParent(successCallback *FileSystemEntryCallback, errorCallback *ErrorCallback) {
	var (
		_args [2]interface{}
		_end  int
	)
	if successCallback != nil {

		var __callback0 js.Value
		if successCallback != nil {
			__callback0 = (*successCallback).Value
		} else {
			__callback0 = js.Null()
		}
		_p0 := __callback0
		_args[0] = _p0
		_end++
	}
	if errorCallback != nil {

		var __callback1 js.Value
		if errorCallback != nil {
			__callback1 = (*errorCallback).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("getParent", _args[0:_end]...)
	return
}

// interface: FileSystemFileEntry
type FileSystemFileEntry struct {
	FileSystemEntry
}

// FileSystemFileEntryFromJS is casting a js.Wrapper into FileSystemFileEntry.
func FileSystemFileEntryFromJS(value js.Wrapper) *FileSystemFileEntry {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileSystemFileEntry{}
	ret.Value_JS = input
	return ret
}

func (_this *FileSystemFileEntry) File(successCallback *FileCallback, errorCallback *ErrorCallback) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if successCallback != nil {
		__callback0 = (*successCallback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if errorCallback != nil {

		var __callback1 js.Value
		if errorCallback != nil {
			__callback1 = (*errorCallback).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("file", _args[0:_end]...)
	return
}
