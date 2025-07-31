package svg

import (
	"github.com/cshabsin/thegrid/js"
)

type Pattern struct {
	js.DOMElement
}

func (p Pattern) Append(child interface{ AsDOM() js.DOMElement }) {
	p.DOMElement.Append(child.AsDOM())
}
