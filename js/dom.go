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
	return DOMElement{document.Get("body"), document}
}

type Attr struct {
	Name  string
	Value interface{}
}

func Class(value string) Attr {
	return Attr{Name: "class", Value: value}
}

func (document DOMDocument) CreateElement(tagName string, attrs ...Attr) DOMElement {
	elem := document.Call("createElement", tagName)
	for _, attr := range attrs {
		elem.Call("setAttribute", attr.Name, attr.Value)
	}
	return DOMElement{elem, document}
}

func (document DOMDocument) CreateElementNS(ns string, tagName string, attrs ...Attr) DOMElement {
	elem := document.Call("createElementNS", ns, tagName)
	for _, attr := range attrs {
		elem.Call("setAttribute", attr.Name, attr.Value)
	}
	return DOMElement{elem, document}
}

func (document DOMDocument) CreateSVG(attrs ...Attr) SVG {
	elem := document.CreateElementNS("http://www.w3.org/2000/svg", "svg", attrs...)
	return SVG(elem)
}

type DOMElement struct {
	js.Value
	document DOMDocument
}

type elementer interface {
	AsDOM() DOMElement
}

func (el DOMElement) AsDOM() DOMElement {
	return el
}

func (el DOMElement) Append(child elementer) {
	el.Call("append", child.AsDOM().Value)
}
