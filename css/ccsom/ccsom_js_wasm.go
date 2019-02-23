// Code generated by webidl-bind. DO NOT EDIT.

package ccsom

import "syscall/js"

// using following types:

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

// interface: CSSGroupingRule
type CSSGroupingRule struct {
	CSSRule
}

// CSSGroupingRuleFromJS is casting a js.Wrapper into CSSGroupingRule.
func CSSGroupingRuleFromJS(value js.Wrapper) *CSSGroupingRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSGroupingRule{}
	ret.Value_JS = input
	return ret
}

// CssRules returning attribute 'cssRules' with
// type CSSRuleList (idl: CSSRuleList).
func (_this *CSSGroupingRule) CssRules() *CSSRuleList {
	var ret *CSSRuleList
	value := _this.Value_JS.Get("cssRules")
	ret = CSSRuleListFromJS(value)
	return ret
}

func (_this *CSSGroupingRule) InsertRule(rule string, index *uint) (_result uint) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := rule
	_args[0] = _p0
	_end++
	if index != nil {
		_p1 := index
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("insertRule", _args[0:_end]...)
	var (
		_converted uint // javascript: unsigned long _what_return_name
	)
	_converted = (uint)((_returned).Int())
	_result = _converted
	return
}

func (_this *CSSGroupingRule) DeleteRule(index uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("deleteRule", _args[0:_end]...)
	return
}

// interface: CSSImportRule
type CSSImportRule struct {
	CSSRule
}

// CSSImportRuleFromJS is casting a js.Wrapper into CSSImportRule.
func CSSImportRuleFromJS(value js.Wrapper) *CSSImportRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSImportRule{}
	ret.Value_JS = input
	return ret
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *CSSImportRule) Href() string {
	var ret string
	value := _this.Value_JS.Get("href")
	ret = (value).String()
	return ret
}

// Media returning attribute 'media' with
// type MediaList (idl: MediaList).
func (_this *CSSImportRule) Media() *MediaList {
	var ret *MediaList
	value := _this.Value_JS.Get("media")
	ret = MediaListFromJS(value)
	return ret
}

// StyleSheet returning attribute 'styleSheet' with
// type CSSStyleSheet (idl: CSSStyleSheet).
func (_this *CSSImportRule) StyleSheet() *CSSStyleSheet {
	var ret *CSSStyleSheet
	value := _this.Value_JS.Get("styleSheet")
	ret = CSSStyleSheetFromJS(value)
	return ret
}

// interface: CSSMarginRule
type CSSMarginRule struct {
	CSSRule
}

// CSSMarginRuleFromJS is casting a js.Wrapper into CSSMarginRule.
func CSSMarginRuleFromJS(value js.Wrapper) *CSSMarginRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSMarginRule{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *CSSMarginRule) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Style returning attribute 'style' with
// type CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSMarginRule) Style() *CSSStyleDeclaration {
	var ret *CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = CSSStyleDeclarationFromJS(value)
	return ret
}

// interface: CSSNamespaceRule
type CSSNamespaceRule struct {
	CSSRule
}

// CSSNamespaceRuleFromJS is casting a js.Wrapper into CSSNamespaceRule.
func CSSNamespaceRuleFromJS(value js.Wrapper) *CSSNamespaceRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSNamespaceRule{}
	ret.Value_JS = input
	return ret
}

// NamespaceURI returning attribute 'namespaceURI' with
// type string (idl: DOMString).
func (_this *CSSNamespaceRule) NamespaceURI() string {
	var ret string
	value := _this.Value_JS.Get("namespaceURI")
	ret = (value).String()
	return ret
}

// Prefix returning attribute 'prefix' with
// type string (idl: DOMString).
func (_this *CSSNamespaceRule) Prefix() string {
	var ret string
	value := _this.Value_JS.Get("prefix")
	ret = (value).String()
	return ret
}

// interface: CSSPageRule
type CSSPageRule struct {
	CSSGroupingRule
}

// CSSPageRuleFromJS is casting a js.Wrapper into CSSPageRule.
func CSSPageRuleFromJS(value js.Wrapper) *CSSPageRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSPageRule{}
	ret.Value_JS = input
	return ret
}

// SelectorText returning attribute 'selectorText' with
// type string (idl: DOMString).
func (_this *CSSPageRule) SelectorText() string {
	var ret string
	value := _this.Value_JS.Get("selectorText")
	ret = (value).String()
	return ret
}

// SetSelectorText setting attribute 'selectorText' with
// type string (idl: DOMString).
func (_this *CSSPageRule) SetSelectorText(value string) {
	input := value
	_this.Value_JS.Set("selectorText", input)
}

// Style returning attribute 'style' with
// type CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSPageRule) Style() *CSSStyleDeclaration {
	var ret *CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = CSSStyleDeclarationFromJS(value)
	return ret
}

// interface: CSSRule
type CSSRule struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSRule) JSValue() js.Value {
	return _this.Value_JS
}

// CSSRuleFromJS is casting a js.Wrapper into CSSRule.
func CSSRuleFromJS(value js.Wrapper) *CSSRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSRule{}
	ret.Value_JS = input
	return ret
}

const STYLERULE_CSSRule int = 1
const CHARSETRULE_CSSRule int = 2
const IMPORTRULE_CSSRule int = 3
const MEDIARULE_CSSRule int = 4
const FONTFACERULE_CSSRule int = 5
const PAGERULE_CSSRule int = 6
const MARGINRULE_CSSRule int = 9
const NAMESPACERULE_CSSRule int = 10

// Type returning attribute 'type' with
// type int (idl: unsigned short).
func (_this *CSSRule) Type() int {
	var ret int
	value := _this.Value_JS.Get("type")
	ret = (value).Int()
	return ret
}

// CssText returning attribute 'cssText' with
// type string (idl: DOMString).
func (_this *CSSRule) CssText() string {
	var ret string
	value := _this.Value_JS.Get("cssText")
	ret = (value).String()
	return ret
}

// SetCssText setting attribute 'cssText' with
// type string (idl: DOMString).
func (_this *CSSRule) SetCssText(value string) {
	input := value
	_this.Value_JS.Set("cssText", input)
}

// ParentRule returning attribute 'parentRule' with
// type CSSRule (idl: CSSRule).
func (_this *CSSRule) ParentRule() *CSSRule {
	var ret *CSSRule
	value := _this.Value_JS.Get("parentRule")
	if value.Type() != js.TypeNull {
		ret = CSSRuleFromJS(value)
	}
	return ret
}

// ParentStyleSheet returning attribute 'parentStyleSheet' with
// type CSSStyleSheet (idl: CSSStyleSheet).
func (_this *CSSRule) ParentStyleSheet() *CSSStyleSheet {
	var ret *CSSStyleSheet
	value := _this.Value_JS.Get("parentStyleSheet")
	if value.Type() != js.TypeNull {
		ret = CSSStyleSheetFromJS(value)
	}
	return ret
}

// interface: CSSRuleList
type CSSRuleList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSRuleList) JSValue() js.Value {
	return _this.Value_JS
}

// CSSRuleListFromJS is casting a js.Wrapper into CSSRuleList.
func CSSRuleListFromJS(value js.Wrapper) *CSSRuleList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSRuleList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *CSSRuleList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *CSSRuleList) Item(index uint) (_result *CSSRule) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *CSSRule // javascript: CSSRule _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = CSSRuleFromJS(_returned)
	}
	_result = _converted
	return
}

// interface: CSSStyleDeclaration
type CSSStyleDeclaration struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSStyleDeclaration) JSValue() js.Value {
	return _this.Value_JS
}

// CSSStyleDeclarationFromJS is casting a js.Wrapper into CSSStyleDeclaration.
func CSSStyleDeclarationFromJS(value js.Wrapper) *CSSStyleDeclaration {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSStyleDeclaration{}
	ret.Value_JS = input
	return ret
}

// CssText returning attribute 'cssText' with
// type string (idl: DOMString).
func (_this *CSSStyleDeclaration) CssText() string {
	var ret string
	value := _this.Value_JS.Get("cssText")
	ret = (value).String()
	return ret
}

// SetCssText setting attribute 'cssText' with
// type string (idl: DOMString).
func (_this *CSSStyleDeclaration) SetCssText(value string) {
	input := value
	_this.Value_JS.Set("cssText", input)
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *CSSStyleDeclaration) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

// ParentRule returning attribute 'parentRule' with
// type CSSRule (idl: CSSRule).
func (_this *CSSStyleDeclaration) ParentRule() *CSSRule {
	var ret *CSSRule
	value := _this.Value_JS.Get("parentRule")
	if value.Type() != js.TypeNull {
		ret = CSSRuleFromJS(value)
	}
	return ret
}

// CssFloat returning attribute 'cssFloat' with
// type string (idl: DOMString).
func (_this *CSSStyleDeclaration) CssFloat() string {
	var ret string
	value := _this.Value_JS.Get("cssFloat")
	ret = (value).String()
	return ret
}

// SetCssFloat setting attribute 'cssFloat' with
// type string (idl: DOMString).
func (_this *CSSStyleDeclaration) SetCssFloat(value string) {
	input := value
	_this.Value_JS.Set("cssFloat", input)
}

func (_this *CSSStyleDeclaration) Item(index uint) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *CSSStyleDeclaration) GetPropertyValue(property string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := property
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getPropertyValue", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *CSSStyleDeclaration) GetPropertyPriority(property string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := property
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getPropertyPriority", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *CSSStyleDeclaration) SetProperty(property string, value string, priority *string) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := property
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	if priority != nil {
		_p2 := priority
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("setProperty", _args[0:_end]...)
	return
}

func (_this *CSSStyleDeclaration) RemoveProperty(property string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := property
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("removeProperty", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// interface: CSSStyleRule
type CSSStyleRule struct {
	CSSRule
}

// CSSStyleRuleFromJS is casting a js.Wrapper into CSSStyleRule.
func CSSStyleRuleFromJS(value js.Wrapper) *CSSStyleRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSStyleRule{}
	ret.Value_JS = input
	return ret
}

// SelectorText returning attribute 'selectorText' with
// type string (idl: DOMString).
func (_this *CSSStyleRule) SelectorText() string {
	var ret string
	value := _this.Value_JS.Get("selectorText")
	ret = (value).String()
	return ret
}

// SetSelectorText setting attribute 'selectorText' with
// type string (idl: DOMString).
func (_this *CSSStyleRule) SetSelectorText(value string) {
	input := value
	_this.Value_JS.Set("selectorText", input)
}

// Style returning attribute 'style' with
// type CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSStyleRule) Style() *CSSStyleDeclaration {
	var ret *CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = CSSStyleDeclarationFromJS(value)
	return ret
}

// interface: CSSStyleSheet
type CSSStyleSheet struct {
	StyleSheet
}

// CSSStyleSheetFromJS is casting a js.Wrapper into CSSStyleSheet.
func CSSStyleSheetFromJS(value js.Wrapper) *CSSStyleSheet {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSStyleSheet{}
	ret.Value_JS = input
	return ret
}

// OwnerRule returning attribute 'ownerRule' with
// type CSSRule (idl: CSSRule).
func (_this *CSSStyleSheet) OwnerRule() *CSSRule {
	var ret *CSSRule
	value := _this.Value_JS.Get("ownerRule")
	if value.Type() != js.TypeNull {
		ret = CSSRuleFromJS(value)
	}
	return ret
}

// CssRules returning attribute 'cssRules' with
// type CSSRuleList (idl: CSSRuleList).
func (_this *CSSStyleSheet) CssRules() *CSSRuleList {
	var ret *CSSRuleList
	value := _this.Value_JS.Get("cssRules")
	ret = CSSRuleListFromJS(value)
	return ret
}

func (_this *CSSStyleSheet) InsertRule(rule string, index *uint) (_result uint) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := rule
	_args[0] = _p0
	_end++
	if index != nil {
		_p1 := index
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("insertRule", _args[0:_end]...)
	var (
		_converted uint // javascript: unsigned long _what_return_name
	)
	_converted = (uint)((_returned).Int())
	_result = _converted
	return
}

func (_this *CSSStyleSheet) DeleteRule(index uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("deleteRule", _args[0:_end]...)
	return
}

// interface: MediaList
type MediaList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaList) JSValue() js.Value {
	return _this.Value_JS
}

// MediaListFromJS is casting a js.Wrapper into MediaList.
func MediaListFromJS(value js.Wrapper) *MediaList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaList{}
	ret.Value_JS = input
	return ret
}

// MediaText returning attribute 'mediaText' with
// type string (idl: DOMString).
func (_this *MediaList) MediaText() string {
	var ret string
	value := _this.Value_JS.Get("mediaText")
	ret = (value).String()
	return ret
}

// SetMediaText setting attribute 'mediaText' with
// type string (idl: DOMString).
func (_this *MediaList) SetMediaText(value string) {
	input := value
	_this.Value_JS.Set("mediaText", input)
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *MediaList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *MediaList) Item(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *MediaList) AppendMedium(medium string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := medium
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("appendMedium", _args[0:_end]...)
	return
}

func (_this *MediaList) DeleteMedium(medium string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := medium
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("deleteMedium", _args[0:_end]...)
	return
}

// interface: StyleSheet
type StyleSheet struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *StyleSheet) JSValue() js.Value {
	return _this.Value_JS
}

// StyleSheetFromJS is casting a js.Wrapper into StyleSheet.
func StyleSheetFromJS(value js.Wrapper) *StyleSheet {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &StyleSheet{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *StyleSheet) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *StyleSheet) Href() *string {
	var ret *string
	value := _this.Value_JS.Get("href")
	if value.Type() != js.TypeNull {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// OwnerNode returning attribute 'ownerNode' with
// type Union (idl: Union).
func (_this *StyleSheet) OwnerNode() *Union {
	var ret *Union
	value := _this.Value_JS.Get("ownerNode")
	if value.Type() != js.TypeNull {
		ret = UnionFromJS(value)
	}
	return ret
}

// ParentStyleSheet returning attribute 'parentStyleSheet' with
// type CSSStyleSheet (idl: CSSStyleSheet).
func (_this *StyleSheet) ParentStyleSheet() *CSSStyleSheet {
	var ret *CSSStyleSheet
	value := _this.Value_JS.Get("parentStyleSheet")
	if value.Type() != js.TypeNull {
		ret = CSSStyleSheetFromJS(value)
	}
	return ret
}

// Title returning attribute 'title' with
// type string (idl: DOMString).
func (_this *StyleSheet) Title() *string {
	var ret *string
	value := _this.Value_JS.Get("title")
	if value.Type() != js.TypeNull {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// Media returning attribute 'media' with
// type MediaList (idl: MediaList).
func (_this *StyleSheet) Media() *MediaList {
	var ret *MediaList
	value := _this.Value_JS.Get("media")
	ret = MediaListFromJS(value)
	return ret
}

// Disabled returning attribute 'disabled' with
// type bool (idl: boolean).
func (_this *StyleSheet) Disabled() bool {
	var ret bool
	value := _this.Value_JS.Get("disabled")
	ret = (value).Bool()
	return ret
}

// SetDisabled setting attribute 'disabled' with
// type bool (idl: boolean).
func (_this *StyleSheet) SetDisabled(value bool) {
	input := value
	_this.Value_JS.Set("disabled", input)
}

// interface: StyleSheetList
type StyleSheetList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *StyleSheetList) JSValue() js.Value {
	return _this.Value_JS
}

// StyleSheetListFromJS is casting a js.Wrapper into StyleSheetList.
func StyleSheetListFromJS(value js.Wrapper) *StyleSheetList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &StyleSheetList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *StyleSheetList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *StyleSheetList) Item(index uint) (_result *CSSStyleSheet) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *CSSStyleSheet // javascript: CSSStyleSheet _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = CSSStyleSheetFromJS(_returned)
	}
	_result = _converted
	return
}
