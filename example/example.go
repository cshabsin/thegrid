package main

import (
	"fmt"
	"sync"
	syscalljs "syscall/js"

	"github.com/cshabsin/thegrid/example/view"
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	hm := hexmap.NewHexMap(10, 11, 70, false)
	mapGroup := document.CreateSVG("g", js.MakeAttr("class", "map-anchor-group"), js.MakeAttr("transform", "translate(10,10)"))
	mapGroup.Append(hm.GridMesh().MakeSVG(document, js.Class("map-mesh")))
	parsec := view.NewParsec(document, hm, 2, 2)
	mapGroup.Append(parsec.Anchor)
	parsec.Anchor.AddEventListener("mouseenter", func(js.DOMElement, syscalljs.Value) {
		parsec.Hilite(true)
	})
	parsec.Anchor.AddEventListener("mouseleave", func(js.DOMElement, syscalljs.Value) {
		parsec.Hilite(false)
	})

	svg := document.GetElementByID("map-svg")
	svg.SetAttr("height", fmt.Sprintf("%fpx", hm.GetPixHeight()+20))
	svg.SetAttr("width", fmt.Sprintf("%fpx", hm.GetPixWidth()+20))
	svg.Append(mapGroup)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
