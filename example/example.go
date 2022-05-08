package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("hi") // console.log
	document := js.Global().Get("document")
	document.Call("write", "hello, world!")
}
