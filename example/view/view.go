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
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}

func (p Parsec) Hilite(v bool) {
	cls := "map-hexagon"
	if v {
		cls = "map-hexagon-hilite"
	}
	p.hexagon.SetAttr("class", cls)
}
