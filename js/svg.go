package js

import "syscall/js"

type SVG struct {
	js.Value
}

func (s SVG) AsDOM() DOMElement {
	return DOMElement(s)
}
