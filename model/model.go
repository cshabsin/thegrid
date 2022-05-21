package model

type MapData struct {
	FirstRow, FirstCol int // offsets for viewing
	HexGrid            [][]*SystemData
}

type SystemData struct {
	name           string
	relRow, relCol int
	Href           string
	description    string
	SuppressPlanet bool
}
