package svg

import (
	"fmt"

	"github.com/cshabsin/thegrid/js/attr"
)

type Path string

func (p Path) ToElement(svg SVG, attrs ...attr.Attr) Element {
	return svg.CreateElement("path", append(attrs, attr.Attr{Name: "d", Value: string(p)})...)
}

func (p Path) MoveAbs(pt Coord, drawn bool) Path {
	cmd := "M"
	if drawn {
		cmd = "L"
	}
	return Path(string(p) + fmt.Sprintf("%s%f,%f\n", cmd, pt.X, pt.Y))
}

func (p Path) MoveRel(vec Vector, drawn bool) Path {
	cmd := "m"
	if drawn {
		cmd = "l"
	}
	return Path(string(p) + fmt.Sprintf("%s%f,%f\n", cmd, vec.DX, vec.DY))
}
