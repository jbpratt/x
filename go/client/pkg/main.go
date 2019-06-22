// +build js,wasm
package main

import "syscall/js"

var doc js.Value

func init() {
	doc = js.Global().Get("document")
}

func main() {
	div := doc.Call("getElementById", "target")

	node := doc.Call("createElement", "div")
	node.Set("innerText", "Hello World")

	div.Call("appendChild", node)
}
