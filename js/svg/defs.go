package svg

import (
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
)

type Defs struct {
	js.DOMElement
	document js.DOMDocument
}

func (svg SVG) Defs() Defs {
	defsEl := svg.document.QuerySelector("defs")
	if defsEl.IsNull() {
		defsEl = svg.document.CreateElementNS("http://www.w3.org/2000/svg", "defs")
		svg.Append(defsEl)
	}
	return Defs{DOMElement: defsEl, document: svg.document}
}

func (d Defs) CreatePattern(attrs ...attr.Attr) Pattern {
	p := Pattern{
		DOMElement: d.document.CreateElementNS("http://www.w3.org/2000/svg", "pattern", attrs...),
	}
	d.Append(p)
	return p
}
