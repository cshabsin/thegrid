package js

import (
	"errors"
	"net/url"
	"syscall/js"

	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/style"
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

func (document DOMDocument) ReadyState() string {
	return document.Get("readyState").String()
}

func (document DOMDocument) Body() DOMElement {
	bodyVal := document.Get("body")
	return DOMElement{bodyVal, document, Style{bodyVal.Get("style")}}
}

func (document DOMDocument) CreateElement(tagName string, attrs ...attr.Attr) DOMElement {
	elemVal := document.Call("createElement", tagName)
	elem := DOMElement{elemVal, document, Style{elemVal.Get("style")}}
	for _, attr := range attrs {
		elem.SetAttr(attr.Name, attr.Value)
	}
	return elem
}

func (document DOMDocument) CreateElementNS(ns string, tagName string, attrs ...attr.Attr) DOMElement {
	elemVal := document.Call("createElementNS", ns, tagName)
	elem := DOMElement{elemVal, document, Style{elemVal.Get("style")}}
	for _, attr := range attrs {
		elem.SetAttr(attr.Name, attr.Value)
	}
	return elem
}

func (document DOMDocument) GetElementByID(id string) DOMElement {
	elemVal := document.Call("getElementById", id)
	return DOMElement{elemVal, document, Style{elemVal.Get("style")}}
}

func (document DOMDocument) AddEventListener(eventName string, fn func(el DOMElement, e DOMEvent)) {
	document.Call("addEventListener", eventName, js.FuncOf(
		func(this js.Value, args []js.Value) any {
			fn(DOMElement{this, document, Style{this.Get("style")}}, DOMEvent{args[0]})
			return nil
		}))
}

type DOMElement struct {
	js.Value
	document DOMDocument
	style    Style
}

func (el DOMElement) IsNull() bool {
	return el.Value.IsNull()
}

func (el DOMElement) Equal(other DOMElement) bool {
	return el.Value.Equal(other.Value)
}

func (el DOMElement) Style() Style {
	return el.style
}

func (el DOMElement) SetStyle(styles ...style.Style) DOMElement {
	for _, s := range styles {
		el.style.Set(s.Name, s.Value)
	}
	return el
}

func (el DOMElement) ClearStyles(styles ...string) DOMElement {
	for _, s := range styles {
		el.style.Set(s, "")
	}
	return el
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

func (el DOMElement) Clear() {
	el.Call("replaceChildren")
}

func (el DOMElement) Remove() {
	el.Call("remove")
}

func (el DOMElement) SetAttr(name string, value any) {
	el.Call("setAttribute", name, value)
}

func (el DOMElement) AddClass(className string) DOMElement {
	el.Get("classList").Call("add", className)
	return el
}

func (el DOMElement) RemoveClass(className string) DOMElement {
	el.Get("classList").Call("remove", className)
	return el
}

func (el DOMElement) SetText(text string) DOMElement {
	el.Set("textContent", text)
	return el
}

func (el DOMElement) AddEventListener(eventName string, fn func(el DOMElement, e DOMEvent)) {
	el.Call("addEventListener", eventName, js.FuncOf(
		func(this js.Value, args []js.Value) any {
			fn(el, DOMEvent{args[0]})
			return nil
		}))
}

func (el DOMElement) QuerySelectorAll(selector string) []DOMElement {
	nodes := el.Call("querySelectorAll", selector)
	var elements []DOMElement
	for i := 0; i < nodes.Length(); i++ {
		node := nodes.Index(i)
		elements = append(elements, DOMElement{node, el.document, Style{node.Get("style")}})
	}
	return elements
}

func (document DOMDocument) QuerySelector(selector string) DOMElement {
	elemVal := document.Call("querySelector", selector)
	if elemVal.IsNull() {
		return DOMElement{Value: js.Null()}
	}
	return DOMElement{elemVal, document, Style{elemVal.Get("style")}}
}

func (el DOMElement) GetBoundingClientRect() js.Value {
	return el.Call("getBoundingClientRect")
}

func (el DOMElement) GetContext(contextType string) js.Value {
	return el.Call("getContext", contextType)
}

type DOMEvent struct {
	js.Value
}

func (el DOMEvent) GetEventType() string {
	return el.Value.Get("type").String()
}

func (el DOMEvent) Key() string {
	return el.Value.Get("key").String()
}

func RequestAnimationFrame(fn func(timestamp float64)) {
	js.Global().Call("requestAnimationFrame", js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(args[0].Float())
		return nil
	}))
}

func Global() js.Value {
	return js.Global()
}

func Confirm(message string) bool {
	return js.Global().Call("confirm", message).Bool()
}

func Null() DOMElement {
	return DOMElement{Value: js.Null()}
}
