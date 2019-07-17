# WebAPI

Go Language Web Assembly bindings for DOM, HTML etc

[![GoDoc](https://godoc.org/github.com/gowebapi/webapi?status.svg)](https://godoc.org/github.com/gowebapi/webapi)

> __WARNING__: The current API is in very early state and should be consider to be expremental. There is upcommig changed like moving types into multiple packages.

This library is trying to be feature complete and up to date with current standard to let everything a browser having to offect be available in WebAssembly in a Go API accessable. To achive this most
of the code is [generated](https://github.com/gowebapi/webidlgenerator) from [WebIDL files](https://github.com/gowebapi/idl). WebIDL files can be found in different standard making it easiser to be up to date than handwritten binding code.

Example:

```golang
func main() {
    element := webapiall.GetWindow().Document().GetElementById("foo")
    // NOTE: After Go 1.12, the extra call to JSValue() will no loger   be needed
    button := webapiall.HTMLButtonElementFromJS(element.JSValue())
    button.SetValue("Hello World")
}
```

## Roadmap

|Milestone|Status|Description|
|----|----|----|
|1|Ongoing|Be abto to read and and understand all of WebIDL syntax. Notible missing thing are _namespace_ and _union_|
|2|Planned|Easy up the syntax. Promise will be refactor and other places where simplifications can be done. Classes can be renamed like clipboard.ClipboardPermission to clipboard.Permission.|
|3|Planned|API Stability with a predefined release schedule. Documentation and added examples|

More infomation of missing stuff can be found on the [generator status info](https://github.com/gowebapi/webidlgenerator).
