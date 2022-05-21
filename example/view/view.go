package view

import (
	"fmt"

	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

type Parsec struct {
	Anchor  js.DOMElement
	hexagon js.DOMElement
}

func NewParsec(document js.DOMDocument, hm *hexmap.HexMap, col, row int) *Parsec {
	hexAnchor := document.CreateSVG("a", js.MakeAttr("class", "map-anchor"), hm.CellTranslate(col, row))
	hexagon := hm.HexPath(document, col, row, "map-hexagon")
	hexAnchor.Append(hexagon) // TODO: just hexagon, but tweak the class with events
	t := document.CreateSVG("text", js.MakeAttr("y", 20), js.MakeAttr("class", "map-name"))
	t.Value.Set("textContent", fmt.Sprintf("%d, %d", col, row))
	hexAnchor.Append(t)
	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon-hilite")
	})
	hexAnchor.AddEventListener("mouseleave", func(_ js.DOMElement, ev js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon")
	})
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}
