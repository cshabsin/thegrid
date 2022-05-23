package svg

import (
	"github.com/cshabsin/thegrid/js"
)

type Coord struct {
	X, Y float64
}

type Vector struct {
	DX, DY float64
}

type SVG struct {
	js.DOMElement

	document js.DOMDocument
}

func MakeSVG(document js.DOMDocument, attrs ...js.Attr) SVG {
	return SVG{
		DOMElement: document.CreateElementNS("http://www.w3.org/2000/svg",
			"svg", attrs...),
		document: document,
	}
}

func GetSVGById(document js.DOMDocument, id string) SVG {
	return SVG{
		DOMElement: document.GetElementByID(id),
		document:   document,
	}
}

func (svg SVG) CreateElement(tagName string, attrs ...js.Attr) Element {
	return Element{
		DOMElement: svg.document.CreateElementNS("http://www.w3.org/2000/svg", tagName, attrs...),
		document:   svg.document,
	}
}

func (svg SVG) Circle(radius float64, attrs ...js.Attr) Element {
	return svg.CreateElement("circle", append(attrs, js.MakeAttr("r", radius))...)
}

func (svg SVG) Text(text string, attrs ...js.Attr) Element {
	elem := svg.CreateElement("text", attrs...)
	elem.Set("textContent", text)
	return elem
}

type Element struct {
	js.DOMElement

	document js.DOMDocument
}
