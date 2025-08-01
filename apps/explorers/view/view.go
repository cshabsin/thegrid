package view

import (
	"fmt"
	"math/rand"

	"github.com/cshabsin/thegrid/apps/explorers/data"
	"github.com/cshabsin/thegrid/apps/explorers/model"
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/svg"
)

type MapView struct {
	SVG         svg.SVG
	HexMap      *hexmap.HexMap
	DataElement js.DOMElement
	Highlighter svg.Element
}

// CreateHexAnchor creates the interactive anchor group for a single hex, but does not append it to the map.
func (mv *MapView) CreateHexAnchor(col, row int, e model.Entity) svg.Element {
	hex := mv.HexMap.Grid[col][row]
	// The anchor is the interactive element. It gets positioned at the hex center.
	hexAnchor := mv.SVG.CreateElement("a", attr.Class("map-anchor"), attr.Translate(hex.CenterX, hex.CenterY))

	// The visible hexagon is a child of the anchor. It has no position itself.
	hexAnchor.Append(hex.ToElement(mv.SVG, mv.HexMap.Radius()))

	if e != nil {
		hexAnchor.Append(mv.SVG.Text(e.Label(), attr.Y(50), attr.Class("map-coord")))
		hexAnchor.Append(mv.SVG.Text(e.Name(), attr.Y(20), attr.Class("map-name")))

		if e.HasCircle() {
			hexAnchor.Append(mv.SVG.Circle(5, attr.Class("map-planet")))
		}
	}

	hexAnchor.AddEventListener("mouseenter", func(js.DOMElement, js.DOMEvent) {
		mv.Highlighter.SetAttr("transform", fmt.Sprintf("translate(%f, %f)", hex.CenterX, hex.CenterY))
		mv.Highlighter.SetAttr("visibility", "visible")
		if e != nil {
			mv.DataElement.Set("innerHTML", e.Description())
		}
	})
	hexAnchor.AddEventListener("mouseleave", func(js.DOMElement, js.DOMEvent) {
		mv.Highlighter.SetAttr("visibility", "hidden")
		if e != nil {
			mv.DataElement.Set("innerHTML", "")
		}
	})

	return hexAnchor
}

func (mv *MapView) NewPathSegment(seg data.PathSegment, cls string, attrs ...attr.Attr) svg.Element {
	var g svg.G
	var p svg.Path
	p.MoveAbs(mv.HexMap.CellCenter(seg.StartCoord[0], seg.StartCoord[1]).Translate(float64(seg.StartOffset[0]), float64(seg.StartOffset[1])), false)
	p.MoveAbs(mv.HexMap.CellCenter(seg.EndCoord[0], seg.EndCoord[1]).Translate(float64(seg.EndOffset[0]), float64(seg.EndOffset[1])), true)
	path := p.WithAttrs(append(attrs, attr.Make("marker-end", "url(#triangle)"), attr.Class(cls))...)
	pathWide := p.WithAttrs(attr.Class(cls + "-wide"))
	g.Append(path)
	g.Append(pathWide)
	group := g.ToElement(mv.SVG)
	group.AddEventListener("mouseenter", func(el js.DOMElement, ev js.DOMEvent) {
		path.ToElement(mv.SVG).AddClass(cls + "-hilite")
		mv.DataElement.Set("innerHTML", seg.Description)
	})
	group.AddEventListener("mouseleave", func(el js.DOMElement, ev js.DOMEvent) {
		path.ToElement(mv.SVG).RemoveClass(cls + "-hilite")
		mv.DataElement.Set("innerHTML", "")
	})
	return group
}

func CreateStarfieldPattern(svgEl svg.SVG) {
	defs := svgEl.Defs()
	starPattern := defs.CreatePattern("starfield",
		attr.Make("width", "100"),
		attr.Make("height", "100"),
		attr.Make("patternUnits", "userSpaceOnUse"),
	)

	for i := 0; i < 20; i++ {
		x := rand.Float64() * 100
		y := rand.Float64() * 100

		star := svgEl.Circle(0.5,
			attr.Make("cx", fmt.Sprintf("%f", x)),
			attr.Make("cy", fmt.Sprintf("%f", y)),
			attr.Make("fill", "white"),
		)
		starPattern.Append(star)
	}
}