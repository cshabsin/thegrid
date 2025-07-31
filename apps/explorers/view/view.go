package view

import (
	"github.com/cshabsin/thegrid/apps/explorers/data/data"
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/svg"
	"github.com/cshabsin/thegrid/model"
)

type Parsec struct {
	Anchor  svg.Element
	hexagon svg.Element
}

type MapView struct {
	SVG    svg.SVG
	HexMap *hexmap.HexMap

	DataElement js.DOMElement
}

func (mv *MapView) NewParsec(col, row int, e model.Entity) *Parsec {
	hexAnchor := mv.SVG.CreateElement("a", attr.Class("map-anchor"), mv.HexMap.CellTranslate(col, row))
	hexagon := mv.HexMap.HexPath(mv.SVG, "map-hexagon")
	hexAnchor.Append(hexagon) // TODO: just hexagon, but tweak the class with events
	hexAnchor.Append(mv.SVG.Text(e.Label(), attr.Y(50), attr.Class("map-coord")))
	hexAnchor.Append(mv.SVG.Text(e.Name(), attr.Y(20), attr.Class("map-name")))

	if e.HasCircle() {
		hexAnchor.Append(mv.SVG.Circle(5, attr.Class("map-planet")))
	}

	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon-hilite")
		mv.DataElement.Set("innerHTML", e.Description())
	})
	hexAnchor.AddEventListener("mouseleave", func(js.DOMElement, js.DOMEvent) {
		hexagon.SetAttr("class", "map-hexagon")
		mv.DataElement.Set("innerHTML", "")
	})
	return &Parsec{Anchor: hexAnchor, hexagon: hexagon}
}

func (mv *MapView) NewPathSegment(seg data.PathSegment, cls string, attrs ...attr.Attr) svg.Element {
	var g svg.G
	var p svg.Path
	p.MoveAbs(mv.HexMap.CellCenter(seg.StartCoord[0], seg.StartCoord[1]).Translate(float64(seg.StartOffset[0]), float64(seg.StartOffset[1])), false)
	p.MoveAbs(mv.HexMap.CellCenter(seg.EndCoord[0], seg.EndCoord[1]).Translate(float64(seg.EndOffset[0]), float64(seg.EndOffset[1])), true)
	g.Append(p.WithAttrs(append(attrs, attr.Make("marker-end", "url(#triangle)"), attr.Class(cls))...))
	g.Append(p.WithAttrs(attr.Class(cls + "-wide")))
	return g.ToElement(mv.SVG)
}
