package view

import (
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

type Parsec struct {
	Anchor  js.DOMElement
	hexagon js.DOMElement
}

type Entity interface {
	Name() string
	Label() string
	HasCircle() bool
}

type MapView struct {
	document js.DOMDocument
	hexMap   *hexmap.HexMap
}

func NewMapView(document js.DOMDocument, hexMap *hexmap.HexMap) *MapView {
	return &MapView{document: document, hexMap: hexMap}
}

func (mv *MapView) NewParsec(col, row int, e Entity) *Parsec {
	hexAnchor := mv.document.CreateSVG("a", js.Class("map-anchor"), mv.hexMap.CellTranslate(col, row))
	hexagon := mv.hexMap.HexPath(mv.document, "map-hexagon")
	hexAnchor.Append(hexagon) // TODO: just hexagon, but tweak the class with events

	t := mv.document.CreateSVG("text", js.MakeAttr("y", 50), js.Class("map-coord"))
	t.Value.Set("textContent", e.Label())
	hexAnchor.Append(t)

	t = mv.document.CreateSVG("text", js.MakeAttr("y", 20), js.Class("map-name"))
	t.Value.Set("textContent", e.Name())
	hexAnchor.Append(t)

	if e.HasCircle() {
		hexAnchor.Append(mv.document.CreateSVG("circle", js.MakeAttr("r", 5), js.Class("map-planet")))
	}

	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon-hilite")
	})
	hexAnchor.AddEventListener("mouseleave", func(_ js.DOMElement, ev js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon")
	})
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}
