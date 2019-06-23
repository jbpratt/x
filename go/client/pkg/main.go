// +build js,wasm
// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
package main

import (
	"log"

	"github.com/gopherjs/websocket"
)

/*func test(div js.Value, s string) {
	node := doc.Call("createElement", "p")
	node.Set("innerText", s)
	div.Call("appendChild", node)
}*/

func main() {
	_, err := websocket.Dial("wss://chat.strims.gg/ws")
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
