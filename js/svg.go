package js

import (
	"fmt"
	"syscall/js"
)

const svgNS = "http://www.w3.org/2000/svg"

type SVG struct {
	js.Value
	document DOMDocument
}

func (s SVG) AsDOM() DOMElement {
	return DOMElement(s)
}

func (s SVG) CreatePath(path SVGPath, attrs ...Attr) DOMElement {
	attrs = append(attrs, Attr{Name: "d", Value: string(path)})
	return s.document.CreateElementNS(svgNS, "path", attrs...)
}

func (s SVG) AddPath(path SVGPath, attrs ...Attr) DOMElement {
	pathElem := s.CreatePath(path, attrs...)
	s.AsDOM().Append(pathElem)
	return pathElem
}

func (s SVG) SetAttr(name string, value interface{}) {
	s.Call("setAttribute", name, value)
}

type SVGCoord struct {
	X, Y float64
}

type SVGVector struct {
	DX, DY float64
}

type SVGPath string

func (p SVGPath) MoveAbs(pt SVGCoord, drawn bool) SVGPath {
	cmd := "M"
	if drawn {
		cmd = "L"
	}
	return SVGPath(string(p) + fmt.Sprintf("%s%f,%f\n", cmd, pt.X, pt.Y))
}

func (p SVGPath) MoveRel(vec SVGVector, drawn bool) SVGPath {
	cmd := "m"
	if drawn {
		cmd = "l"
	}
	return SVGPath(string(p) + fmt.Sprintf("%s%f,%f\n", cmd, vec.DX, vec.DY))
}
