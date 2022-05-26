package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cshabsin/thegrid/example/server/data"
	"github.com/cshabsin/thegrid/example/view"
)

func FromURL(url url.URL) (*MapData, error) {
	jsonData, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer jsonData.Body.Close()
	return FromJSON(jsonData.Body)
}

func FromJSON(r io.Reader) (*MapData, error) {
	var md data.MapData
	err := json.NewDecoder(r).Decode(&md)
	if err != nil {
		return nil, err
	}
	numRows := md.MaxRow - md.MinRow + 1
	numCols := md.MaxCol - md.MinCol + 1
	hexGrid := make([][]view.Entity, numCols)
	for col := range hexGrid {
		hexGrid[col] = make([]view.Entity, numRows)
		for row := range hexGrid[col] {
			hexGrid[col][row] = emptySystem{sysCol: col + md.MinCol, sysRow: row + md.MinRow}
		}
	}
	for _, sys := range md.Systems {
		sysData := &systemData{
			name:           sys.Name,
			sysCol:         sys.SysCol,
			sysRow:         sys.SysRow,
			description:    sys.Description,
			suppressPlanet: sys.SuppressPlanet,
		}
		if sys.ShortSystem != "" {
			sysData.href = "http://scripts.mit.edu/~ringrose/explorers/index.php?title=" + sys.ShortSystem
		}
		hexGrid[sys.SysCol-md.MinCol][sys.SysRow-md.MinRow] = sysData
	}
	return &MapData{
		FirstCol: md.MinCol,
		FirstRow: md.MinRow,
		HexGrid:  hexGrid,
	}, nil
}

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
