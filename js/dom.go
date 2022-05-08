package js

import "syscall/js"

type DOMDocument struct {
	js.Value
}

// Document returns the "document" from the global scope.
func Document() DOMDocument {
	return DOMDocument{js.Global().Get("document")}
}

func (document DOMDocument) Body() DOMElement {
	return DOMElement{document.Get("body")}
}

func (document DOMDocument) CreateElement(tagName string) DOMElement {
	return DOMElement{document.Call("createElement", tagName)}
}

type DOMElement struct {
	js.Value
}

func (el DOMElement) Append(child DOMElement) {
	el.Call("append", child.Value)
}
