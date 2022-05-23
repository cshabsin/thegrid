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

func MakeAttr(name string, value interface{}) Attr {
	return Attr{Name: name, Value: value}
}

func Class(value string) Attr {
	return Attr{Name: "class", Value: value}
}

func (document DOMDocument) CreateElement(tagName string, attrs ...Attr) DOMElement {
	elem := DOMElement{document.Call("createElement", tagName), document}
	for _, attr := range attrs {
		elem.SetAttr(attr.Name, attr.Value)
	}
	return elem
}

func (document DOMDocument) CreateElementNS(ns string, tagName string, attrs ...Attr) DOMElement {
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

type DOMEvent struct {
	js.Value
}

func (el DOMEvent) GetEventType() string {
	return el.Value.Get("type").String()
}
