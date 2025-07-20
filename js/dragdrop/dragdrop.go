package dragdrop

import (
	"github.com/cshabsin/thegrid/js"
)

type Draggable struct {
	js.DOMElement
	OnDragStart func(js.DOMEvent)
}

func NewDraggable(el js.DOMElement, onDragStart func(js.DOMEvent)) *Draggable {
	d := &Draggable{el, onDragStart}
	d.AddEventListener("dragstart", func(_ js.DOMElement, e js.DOMEvent) {
		e.Value.Get("dataTransfer").Call("setData", "text/plain", "")
		if d.OnDragStart != nil {
			d.OnDragStart(e)
		}
	})
	return d
}

type DropTarget struct {
	js.DOMElement
	OnDrop      func(js.DOMEvent)
	OnDragOver  func(js.DOMEvent)
	OnDragEnter func(js.DOMEvent)
	OnDragLeave func(js.DOMEvent)
}

func NewDropTarget(el js.DOMElement, onDrop func(js.DOMEvent)) *DropTarget {
	d := &DropTarget{DOMElement: el, OnDrop: onDrop}
	d.AddEventListener("dragover", func(_ js.DOMElement, e js.DOMEvent) {
		e.Value.Call("preventDefault")
		if d.OnDragOver != nil {
			d.OnDragOver(e)
		}
	})
		d.AddEventListener("drop", func(_ js.DOMElement, e js.DOMEvent) {
			e.Value.Call("preventDefault")
		if d.OnDrop != nil {
			d.OnDrop(e)
		}
	})
	d.AddEventListener("dragenter", func(_ js.DOMElement, e js.DOMEvent) {
		if d.OnDragEnter != nil {
			d.OnDragEnter(e)
		}
	})
	d.AddEventListener("dragleave", func(_ js.DOMElement, e js.DOMEvent) {
		if d.OnDragLeave != nil {
			d.OnDragLeave(e)
		}
	})
	return d
}