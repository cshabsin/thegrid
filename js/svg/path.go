package svg

import (
	"fmt"

	"github.com/cshabsin/thegrid/js/attr"
)

type Path struct {
	path  string
	attrs []attr.Attr
}

func (p *Path) ToElement(svg SVG, attrs ...attr.Attr) Element {
	attrs = append(attrs, p.attrs...)
	attrs = append(attrs, attr.Attr{Name: "d", Value: p.path})
	return svg.CreateElement("path", attrs...)
}

func (p *Path) WithAttrs(attrs ...attr.Attr) *Path {
	return &Path{
		path:  p.path,
		attrs: append(p.attrs[:], attrs...),
	}
}
func (p *Path) MoveAbs(pt Coord, drawn bool) {
	cmd := "M"
	if drawn {
		cmd = "L"
	}
	p.path += fmt.Sprintf("%s%f,%f\n", cmd, pt.X, pt.Y)
}

func (p *Path) MoveRel(vec Vector, drawn bool) {
	cmd := "m"
	if drawn {
		cmd = "l"
	}
	p.path += fmt.Sprintf("%s%f,%f\n", cmd, vec.DX, vec.DY)
}

func (p *Path) SetAttr(attrs ...attr.Attr) {
	p.attrs = append(p.attrs, attrs...)
}
