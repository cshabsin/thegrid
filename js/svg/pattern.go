package svg

import (
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
)

type Pattern struct {
	js.DOMElement
}

func NewPattern(doc js.DOMDocument, attrs ...attr.Attr) Pattern {
	return Pattern{
		DOMElement: doc.CreateElementNS("http://www.w3.org/2000/svg", "pattern", attrs...),
	}
}

func (p Pattern) Append(child interface{ AsDOM() js.DOMElement }) {
	p.DOMElement.Append(child.AsDOM())
}
