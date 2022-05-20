package js

import (
	"fmt"
)

type SVGCoord struct {
	X, Y float64
}

type SVGVector struct {
	DX, DY float64
}

type SVGPath string

func (p SVGPath) MakeSVG(document DOMDocument, attrs ...Attr) DOMElement {
	return document.CreateSVG("path", append([]Attr{{Name: "d", Value: string(p)}}, attrs...)...)
}

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
