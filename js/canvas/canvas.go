package canvas

import (
	"github.com/cshabsin/thegrid/js"
)

type Canvas struct {
	js.DOMElement
}

func Get(id string) Canvas {
	return Canvas{js.Document().GetElementByID(id)}
}

func (c Canvas) GetContext(contextType string) Context {
	return Context{js.DOMElement{Value: c.DOMElement.GetContext(contextType)}}
}

func (c Canvas) FullSize() {
	width := c.Get("clientWidth").Float()
	height := c.Get("clientHeight").Float()
	c.SetAttr("width", width)
	c.SetAttr("height", height)
}

type Context struct {
	js.DOMElement
}

func (c Context) BeginPath() {
	c.Call("beginPath")
}

func (c Context) MoveTo(x, y float64) {
	c.Call("moveTo", x, y)
}

func (c Context) LineTo(x, y float64) {
	c.Call("lineTo", x, y)
}

func (c Context) Stroke() {
	c.Call("stroke")
}

func (c Context) ClearRect(x, y, width, height float64) {
	c.Call("clearRect", x, y, width, height)
}

func (c Context) SetStrokeStyle(style string) {
	c.Set("strokeStyle", style)
}

func (c Context) SetLineWidth(width int) {
	c.Set("lineWidth", width)
}

func (c Context) Scale(x, y float64) {
	c.Call("scale", x, y)
}
