// +build js,wasm
package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var doc js.Value

func init() {
	doc = js.Global().Get("document")
}

func test(done chan string, div js.Value) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		node := doc.Call("createElement", "p")
		node.Set("innerText", i)
		div.Call("appendChild", node)
	}
	done <- "goroutine finished"
}

func main() {
	div := doc.Call("getElementById", "target")

	done := make(chan string)
	go test(done, div)

	msg := <-done
	fmt.Println(msg)
}
