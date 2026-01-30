package svg

import (
	"github.com/cshabsin/thegrid/js/attr"
)

type elementer interface {
	ToElement(svg SVG, attrs ...attr.Attr) Element
}

type G []elementer

func (g *G) ToElement(svg SVG, attrs ...attr.Attr) Element {
	gel := svg.CreateElement("g", attrs...)
	for _, elem := range *g {
		gel.Append(elem.ToElement(svg))
	}
	return gel
}

func (g *G) Append(child elementer) {
	*g = append(*g, child)
}
