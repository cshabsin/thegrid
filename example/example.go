package main

import (
	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	p := document.CreateElement("p")
	p.Set("innerHTML", "hi")
	document.Body().Append(p)
}
