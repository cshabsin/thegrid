package svg

import (
	"github.com/cshabsin/thegrid/js/attr"
)

type elementer interface {
	ToElement(svg SVG, attrs ...attr.Attr) Element
}

type G []elementer

func (p G) ToElement(svg SVG, attrs ...attr.Attr) Element {
	gel := svg.CreateElement("g", attrs...)
	for _, elem := range p {
		gel.Append(elem.ToElement(svg))
	}
	return gel
}

func (p G) Append(elem elementer) G {
	return append(p, elem)
}
