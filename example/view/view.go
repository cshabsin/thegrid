package view

import (
	"fmt"

	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
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
	Description() string
}

type MapView struct {
	SVG    svg.SVG
	HexMap *hexmap.HexMap

	DataElement js.DOMElement
}

func (mv *MapView) NewParsec(col, row int, e Entity) *Parsec {
	hexAnchor := mv.SVG.CreateElement("a", attr.Class("map-anchor"), mv.HexMap.CellTranslate(col, row))
	hexagon := mv.HexMap.HexPath(mv.SVG, "map-hexagon")
	hexAnchor.Append(hexagon) // TODO: just hexagon, but tweak the class with events
	hexAnchor.Append(mv.SVG.Text(e.Label(), attr.Make("y", 50), attr.Class("map-coord")))
	hexAnchor.Append(mv.SVG.Text(e.Name(), attr.Make("y", 20), attr.Class("map-name")))

	if e.HasCircle() {
		hexAnchor.Append(mv.SVG.Circle(5, attr.Class("map-planet")))
	}

	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		fmt.Println("hi")
		hexagon.SetAttr("class", "map-hexagon-hilite")
		mv.DataElement.Set("innerHTML", e.Description())
	})
	hexAnchor.AddEventListener("mouseleave", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon")
		mv.DataElement.Set("innerHTML", "")
	})
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}
