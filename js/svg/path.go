package svg

import (
	"fmt"

	"github.com/cshabsin/thegrid/js"
)

type Path string

func (p Path) ToElement(svg SVG, attrs ...js.Attr) Element {
	return svg.CreateElement("path", append([]js.Attr{{Name: "d", Value: string(p)}}, attrs...)...)
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
