package view

import (
	"fmt"

	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/svg"
)

type Parsec struct {
	Anchor  svg.Element
	hexagon svg.Element
}

type Entity interface {
	Name() string
	Label() string
	HasCircle() bool
}

type MapView struct {
	svg    svg.SVG
	hexMap *hexmap.HexMap
}

func NewMapView(svg svg.SVG, hexMap *hexmap.HexMap) *MapView {
	return &MapView{svg: svg, hexMap: hexMap}
}

func (mv *MapView) NewParsec(col, row int, e Entity) *Parsec {
	hexAnchor := mv.svg.CreateElement("a", js.Class("map-anchor"), mv.hexMap.CellTranslate(col, row))
	hexagon := mv.hexMap.HexPath(mv.svg, "map-hexagon")
	hexAnchor.Append(hexagon) // TODO: just hexagon, but tweak the class with events
	hexAnchor.Append(mv.svg.Text(e.Label(), js.MakeAttr("y", 50), js.Class("map-coord")))
	hexAnchor.Append(mv.svg.Text(e.Name(), js.MakeAttr("y", 20), js.Class("map-name")))

	if e.HasCircle() {
		hexAnchor.Append(mv.svg.Circle(5, js.Class("map-planet")))
	}

	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		fmt.Println("hi")
		hexagon.SetAttr("class", "map-hexagon-hilite")
	})
	hexAnchor.AddEventListener("mouseleave", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon")
	})
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}
