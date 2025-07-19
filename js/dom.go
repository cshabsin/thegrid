package js

import (
	"errors"
	"net/url"
	"syscall/js"

	"github.com/cshabsin/thegrid/js/attr"
)

func URL() (*url.URL, error) {
	window := js.Global().Get("window")
	if window.IsNull() {
		return nil, errors.New("window not found")
	}
	location := window.Get("location")
	if location.IsNull() {
		return nil, errors.New("location not found")
	}
	href := location.Get("href")
	if href.IsNull() {
		return nil, errors.New("href not found")
	}
	return url.Parse(href.String())
}

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

func (document DOMDocument) CreateElement(tagName string, attrs ...attr.Attr) DOMElement {
	elem := DOMElement{document.Call("createElement", tagName), document}
	for _, attr := range attrs {
		elem.SetAttr(attr.Name, attr.Value)
	}
	return elem
}

func (document DOMDocument) CreateElementNS(ns string, tagName string, attrs ...attr.Attr) DOMElement {
	elem := DOMElement{document.Call("createElementNS", ns, tagName), document}
	for _, attr := range attrs {
		elem.SetAttr(attr.Name, attr.Value)
	}
	return elem
}

func (document DOMDocument) GetElementByID(id string) DOMElement {
	elem := document.Call("getElementById", id)
	return DOMElement{elem, document}
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

func (el DOMElement) Append(child elementer) DOMElement {
	el.Call("append", child.AsDOM().Value)
	return el
}

func (el DOMElement) SetAttr(name string, value interface{}) {
	el.Call("setAttribute", name, value)
}

func (el DOMElement) AddEventListener(eventName string, fn func(el DOMElement, e DOMEvent)) {
	el.Call("addEventListener", eventName, js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			fn(el, DOMEvent{args[0]})
			return nil
		}))
}

func (el DOMElement) QuerySelectorAll(selector string) []DOMElement {
	nodes := el.Call("querySelectorAll", selector)
	var elements []DOMElement
	for i := 0; i < nodes.Length(); i++ {
		elements = append(elements, DOMElement{nodes.Index(i), el.document})
	}
	return elements
}

func (el DOMElement) GetBoundingClientRect() js.Value {
	return el.Call("getBoundingClientRect")
}

type DOMEvent struct {
	js.Value
}

func (el DOMEvent) GetEventType() string {
	return el.Value.Get("type").String()
}
