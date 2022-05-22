package model

import (
	"fmt"

	"github.com/cshabsin/thegrid/example/view"
)

type MapData struct {
	FirstCol, FirstRow int // offsets for viewing
	HexGrid            [][]view.Entity
}

func (md MapData) GetCell(col, row int) view.Entity {
	c := md.HexGrid[col]
	return c[row]
}

type systemData struct {
	name           string
	sysRow, sysCol int
	href           string
	description    string
	suppressPlanet bool
}

func (s systemData) Name() string {
	return s.name
}

func (s systemData) Label() string {
	return fmt.Sprintf("%02d%02d", s.sysCol, s.sysRow)
}

func (s systemData) HasCircle() bool {
	if s.name == "" {
		return false
	}
	if s.suppressPlanet {
		return false
	}
	return true
}

type emptySystem struct {
	sysRow, sysCol int
}

func (s emptySystem) Name() string {
	return ""
}
func (s emptySystem) Label() string {
	return fmt.Sprintf("%02d%02d", s.sysCol, s.sysRow)
}
func (s emptySystem) HasCircle() bool {
	return false
}
