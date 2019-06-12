// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
package main

import (
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {
	js.Global().Set("printMsg", js.FuncOf(printMsg))
	<-c
	println("here")
}

func printMsg(this js.Value, inputs []js.Value) {
	callback := inputs[len(inputs)-1:][0]
	message := inputs[0].String()

	callback.Invoke(js.Null(), "Did you say "+message)
}
