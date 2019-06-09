// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
package main

import (
	"syscall/js"
)

func main() {
	c := make(chan bool)
	js.Global().Set("printMsg", js.FuncOf(printMsg))
	<-c
}

func printMsg(this js.Value, inputs []js.Value) interface{} {
	msg := inputs[0].String()
	doc := js.Global().Get("document")
	p := doc.Call("createElement", "p")
	p.Set("innerHTML", msg)
	doc.Get("body").Call("appendChild", p)
	return nil
}
